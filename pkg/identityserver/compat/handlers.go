// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package compat

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/gogo/protobuf/types"
	"github.com/labstack/echo"
	"go.thethings.network/lorawan-stack/pkg/pfconfig"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

// gatewayInfoResponse represents a minimal gateway info response for The Things Gateway.
type gatewayInfoResponse struct {
	FrequencyPlanID  string `json:"frequency_plan"`
	FrequencyPlanURL string `json:"frequency_plan_url"`
	FirmwareURL      string `json:"firmware_url"`
	Router           struct {
		MQTTAddress string `json:"mqtt_address"`
	} `json:"router"`
	AutoUpdate bool `json:"auto_update"`
}

func (s *Server) handleGatewayInfo(c echo.Context) error {
	gatewayIDs := c.Get(gatewayIDKey).(ttnpb.GatewayIdentifiers)
	if gateway, err := s.gtwStore.GetGateway(s.ctx, &gatewayIDs,
		&types.FieldMask{Paths: []string{
			"frequency_plan_id",
			"gateway_server_address",
			"update_channel",
			"auto_update",
		}}); err != nil {
		return err
	} else {
		freqPlanURL := &url.URL{
			Scheme: c.Scheme(),
			Host:   c.Request().Host,
			Path:   fmt.Sprintf("%v/frequency-plans/%v", compatAPIPrefix, gateway.FrequencyPlanID),
		}
		response := &gatewayInfoResponse{
			FrequencyPlanID:  gateway.FrequencyPlanID,
			FrequencyPlanURL: freqPlanURL.String(),
			FirmwareURL:      adaptUpdateChannel(gateway.UpdateChannel),
			AutoUpdate:       gateway.AutoUpdate,
		}
		response.Router.MQTTAddress, err = adaptGatewayAddress(gateway.GatewayServerAddress)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (s *Server) handleFreqPlanInfo(c echo.Context) error {
	freqPlanID := c.Param(frequencyPlanIDKey)
	plan, err := s.fpStore.GetByID(freqPlanID)
	if err != nil {
		return err
	}
	config, err := pfconfig.BuildSX1301Config(plan)
	if err != nil {
		return err
	}
	config.TxLUTConfigs = config.TxLUTConfigs[:0]
	return c.JSON(http.StatusOK, struct {
		SX1301Conf *pfconfig.SX1301Config `json:"SX1301_conf"`
	}{
		SX1301Conf: config,
	})
}

const defaultFirmwarePath = "https://thethingsproducts.blob.core.windows.net/the-things-gateway/v1"

// adaptUpdateChannel prepends the default firmware path if the channel itself is not an URL.
func adaptUpdateChannel(channel string) string {
	if channel == "" {
		return ""
	}
	if _, err := url.ParseRequestURI(channel); err != nil {
		return fmt.Sprintf("%v/%v", defaultFirmwarePath, channel)
	}
	return channel
}

// adaptGatewayAddress prepends the gateway address with the scheme "mqtts" and appends
// the port "8882" if they have not been mentioned. If the scheme has been given, no
// changes are done to the address.
func adaptGatewayAddress(address string) (result string, err error) {
	if address == "" {
		return address, nil
	}
	if strings.Contains(address, "://") {
		return address, nil
	}
	var host, port string
	if strings.Contains(address, ":") {
		host, port, err = net.SplitHostPort(address)
		if err != nil {
			return "", err
		}
	} else {
		host = address
	}
	if port == "" {
		port = "8882"
	}
	return fmt.Sprintf("mqtts://%v:%v", host, port), nil
}
