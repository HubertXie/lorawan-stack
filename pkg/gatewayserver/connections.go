// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gatewayserver

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"go.thethings.network/lorawan-stack/pkg/errors"
	"go.thethings.network/lorawan-stack/pkg/gatewayserver/scheduling"
	"go.thethings.network/lorawan-stack/pkg/gatewayserver/udp"
	"go.thethings.network/lorawan-stack/pkg/toa"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

const (
	pullDataExpiration = 1 * time.Minute
	// DownlinkTiming before sending to a gateway connected over UDP,
	// and that does not have a JIT queue.
	DownlinkTiming = -1 * 1000 * time.Millisecond
)

type connection interface {
	addUpstreamObservations(*ttnpb.GatewayUp)
	addDownstreamObservations(*ttnpb.GatewayDown)
	getObservations() ttnpb.GatewayObservations

	gateway() *ttnpb.Gateway
	send(*ttnpb.DownlinkMessage) error
	Close() error
}

type connectionData struct {
	observations   ttnpb.GatewayObservations
	observationsMu sync.RWMutex

	scheduler scheduling.Scheduler
}

func (c *connectionData) schedule(down *ttnpb.DownlinkMessage) (err error) {
	span := scheduling.Span{
		Start: scheduling.ConcentratorTime(down.TxMetadata.Timestamp),
	}
	span.Duration, err = toa.Compute(down.RawPayload, down.Settings)
	if err != nil {
		return errors.NewWithCause(err, "Could not compute time-on-air of the downlink")
	}

	err = c.scheduler.ScheduleAt(span, down.Settings.Frequency)
	if err != nil {
		return errors.NewWithCause(err, "Could not schedule downlink")
	}

	return nil
}

func (c *connectionData) getObservations() ttnpb.GatewayObservations {
	c.observationsMu.Lock()
	observations := c.observations
	c.observationsMu.Unlock()
	return observations
}

func (c *connectionData) addUpstreamObservations(up *ttnpb.GatewayUp) {
	now := time.Now().UTC()

	c.observationsMu.Lock()

	if up.GatewayStatus != nil {
		c.observations.LastStatus = up.GatewayStatus
		c.observations.LastStatusReceivedAt = &now
	}
	if len(up.UplinkMessages) != 0 {
		c.observations.LastUplinkReceivedAt = &now
	}

	c.observationsMu.Unlock()
}

func (c *connectionData) addDownstreamObservations(down *ttnpb.GatewayDown) {
	now := time.Now().UTC()
	c.observationsMu.Lock()
	c.observations.LastDownlinkReceivedAt = &now
	c.observationsMu.Unlock()
}

type gRPCConnection struct {
	connectionData

	link   ttnpb.GtwGs_LinkServer
	cancel context.CancelFunc
	gtw    *ttnpb.Gateway
}

func (c *gRPCConnection) send(down *ttnpb.DownlinkMessage) error {
	if err := c.schedule(down); err != nil {
		return err
	}
	return c.link.Send(&ttnpb.GatewayDown{DownlinkMessage: down})
}

func (c *gRPCConnection) gateway() *ttnpb.Gateway {
	return c.gtw
}

func (c *gRPCConnection) Close() error {
	c.cancel()
	return nil
}

type udpConnection struct {
	connectionData

	gtw                 atomic.Value
	lastPullDataStorage atomic.Value
	lastPullDataTime    atomic.Value

	concentratorStart atomic.Value
	hasSentTxAck      atomic.Value
}

func (c *udpConnection) hasJITQueue() bool {
	hasSentTxAck, ok := c.hasSentTxAck.Load().(bool)
	return ok && hasSentTxAck
}

// Takes a timestamp in microseconds
func (c *udpConnection) syncClock(timestamp uint32) {
	start := time.Now().Add(-1 * time.Microsecond * time.Duration(timestamp))
	c.concentratorStart.Store(start)
}

// Takes a timestamp in microseconds
func (c *udpConnection) realTime(timestamp uint32) (time.Time, bool) {
	concentratorStart, ok := c.concentratorStart.Load().(time.Time)
	if !ok {
		return time.Now(), false
	}

	t := concentratorStart.Add(time.Microsecond * time.Duration(timestamp))
	if t.Before(time.Now()) {
		t = t.Add(time.Duration(int64(1<<32) * 1000))
	}
	return t, true
}

func (c *udpConnection) lastPullData() *udp.Packet {
	pkt, _ := c.lastPullDataStorage.Load().(*udp.Packet)
	return pkt
}

func (c *udpConnection) pullDataExpired() bool {
	lastReceived, ok := c.lastPullDataTime.Load().(time.Time)
	if !ok {
		return true
	}
	return time.Since(lastReceived) > pullDataExpiration
}

func (c *udpConnection) send(down *ttnpb.DownlinkMessage) error {
	gtw := c.gateway()
	if c.pullDataExpired() {
		return errors.NewWithCausef(ErrGatewayNotConnected.New(errors.Attributes{
			"gateway_id": gtw.GetGatewayID(),
		}), "No PULL_DATA received in the last %s", pullDataExpiration.String())
	}

	tx, err := udp.TranslateDownstream(down)
	if err != nil {
		return ErrTranslationFromProtobuf.New(nil)
	}
	if err := c.schedule(down); err != nil {
		return err
	}

	pkt := *c.lastPullData()
	pkt.PacketType = udp.PullResp
	pkt.Data = &udp.Data{TxPacket: &tx}

	writePacket := func() error { return pkt.GatewayConn.Write(&pkt) }

	if gtw.DisableTxDelay || c.hasJITQueue() {
		return writePacket()
	}

	realTime, ok := c.realTime(pkt.Data.TxPacket.Tmst)
	if !ok {
		return writePacket()
	}

	<-time.After(time.Until(realTime.Add(DownlinkTiming)))
	return writePacket()
}

func (c *udpConnection) gateway() *ttnpb.Gateway {
	gtw, _ := c.gtw.Load().(*ttnpb.Gateway)
	return gtw
}

func (c *udpConnection) Close() error { return nil }