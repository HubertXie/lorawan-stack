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

package applicationserver_test

import (
	"context"
	"testing"
	"time"

	pbtypes "github.com/gogo/protobuf/types"
	"github.com/mohae/deepcopy"
	"github.com/smartystreets/assertions"
	. "go.thethings.network/lorawan-stack/pkg/applicationserver"
	"go.thethings.network/lorawan-stack/pkg/applicationserver/redis"
	"go.thethings.network/lorawan-stack/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/pkg/component"
	"go.thethings.network/lorawan-stack/pkg/config"
	"go.thethings.network/lorawan-stack/pkg/errors"
	"go.thethings.network/lorawan-stack/pkg/rpcmetadata"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/types"
	"go.thethings.network/lorawan-stack/pkg/unique"
	"go.thethings.network/lorawan-stack/pkg/util/test"
	"go.thethings.network/lorawan-stack/pkg/util/test/assertions/should"
	"google.golang.org/grpc"
)

func TestDeviceRegistry(t *testing.T) {
	a := assertions.New(t)

	ctx := test.Context()
	is, isAddr := startMockIS(ctx)

	// Register the application in the Entity Registry.
	is.add(ctx, registeredApplicationID, registeredApplicationKey)

	redisClient, flush := test.NewRedis(t, "applicationserver_test")
	defer flush()
	defer redisClient.Close()
	deviceRegistry := &redis.DeviceRegistry{Redis: redisClient}

	c := component.MustNew(test.GetLogger(t), &component.Config{
		ServiceBase: config.ServiceBase{
			GRPC: config.GRPC{
				Listen:                      ":9184",
				AllowInsecureForCredentials: true,
			},
			Cluster: config.Cluster{
				IdentityServer: isAddr,
			},
		},
	})
	config := &Config{
		LinkMode: "explicit",
		Devices:  deviceRegistry,
	}
	_, err := New(c, config)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	test.Must(nil, c.Start())
	defer c.Close()
	mustHavePeer(ctx, c, ttnpb.PeerInfo_ENTITY_REGISTRY)

	conn, err := grpc.Dial(":9184", grpc.WithInsecure(), grpc.WithBlock())
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	defer conn.Close()
	creds := grpc.PerRPCCredentials(rpcmetadata.MD{
		AuthType:      "Bearer",
		AuthValue:     registeredApplicationKey,
		AllowInsecure: true,
	})
	client := ttnpb.NewAsEndDeviceRegistryClient(conn)

	// Unauthorized: no credentials.
	{
		_, err := client.Get(ctx, &ttnpb.GetEndDeviceRequest{
			EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
		})
		a.So(errors.IsUnauthenticated(err), should.BeTrue)

		_, err = client.Set(ctx, &ttnpb.SetEndDeviceRequest{
			Device: ttnpb.EndDevice{
				EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
				FrequencyPlanID:      "EU_863_870",
			},
			FieldMask: pbtypes.FieldMask{
				Paths: []string{"frequency_plan_id"},
			},
		})
		a.So(errors.IsUnauthenticated(err), should.BeTrue)

		_, err = client.Delete(ctx, &registeredDevice.EndDeviceIdentifiers)
		a.So(errors.IsUnauthenticated(err), should.BeTrue)
	}

	// Unauthorized: wrong application.
	{
		otherID := ttnpb.EndDeviceIdentifiers{
			ApplicationIdentifiers: ttnpb.ApplicationIdentifiers{
				ApplicationID: "other-app",
			},
			DeviceID: "other-device",
		}

		_, err := client.Get(ctx, &ttnpb.GetEndDeviceRequest{
			EndDeviceIdentifiers: otherID,
		}, creds)
		a.So(errors.IsPermissionDenied(err), should.BeTrue)

		_, err = client.Set(ctx, &ttnpb.SetEndDeviceRequest{
			Device: ttnpb.EndDevice{
				EndDeviceIdentifiers: otherID,
				FrequencyPlanID:      "EU_863_870",
			},
			FieldMask: pbtypes.FieldMask{
				Paths: []string{"frequency_plan_id"},
			},
		}, creds)
		a.So(errors.IsPermissionDenied(err), should.BeTrue)

		_, err = client.Delete(ctx, &otherID, creds)
		a.So(errors.IsPermissionDenied(err), should.BeTrue)
	}

	// Happy flow: create, update and delete.
	{
		// Assert the device doesn't exist yet.
		_, err := client.Get(ctx, &ttnpb.GetEndDeviceRequest{
			EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
			FieldMask: pbtypes.FieldMask{
				Paths: []string{"ids", "version_ids", "formatters"},
			},
		}, creds)
		a.So(errors.IsNotFound(err), should.BeTrue)

		// Create and assert resemblance.
		_, err = client.Set(ctx, &ttnpb.SetEndDeviceRequest{
			Device: *registeredDevice,
			FieldMask: pbtypes.FieldMask{
				Paths: []string{"ids", "version_ids", "formatters"},
			},
		}, creds)
		a.So(err, should.BeNil)
		dev, err := client.Get(ctx, &ttnpb.GetEndDeviceRequest{
			EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
			FieldMask: pbtypes.FieldMask{
				Paths: []string{"ids", "version_ids", "formatters"},
			},
		}, creds)
		a.So(err, should.BeNil)
		registeredDevice.CreatedAt = dev.CreatedAt
		registeredDevice.UpdatedAt = dev.UpdatedAt
		a.So(dev, should.HaveEmptyDiff, registeredDevice)

		// Update and assert new value.
		registeredDevice.Formatters.UpFormatter = ttnpb.PayloadFormatter_FORMATTER_NONE
		_, err = client.Set(ctx, &ttnpb.SetEndDeviceRequest{
			Device: *registeredDevice,
			FieldMask: pbtypes.FieldMask{
				Paths: []string{"formatters"},
			},
		}, creds)
		a.So(err, should.BeNil)
		dev, err = client.Get(ctx, &ttnpb.GetEndDeviceRequest{
			EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
			FieldMask: pbtypes.FieldMask{
				Paths: []string{"ids", "version_ids", "formatters"},
			},
		}, creds)
		a.So(err, should.BeNil)
		registeredDevice.CreatedAt = dev.CreatedAt
		registeredDevice.UpdatedAt = dev.UpdatedAt
		a.So(dev, should.HaveEmptyDiff, registeredDevice)

		// Delete and assert it's gone.
		_, err = client.Delete(ctx, &registeredDevice.EndDeviceIdentifiers, creds)
		a.So(err, should.BeNil)
		_, err = client.Get(ctx, &ttnpb.GetEndDeviceRequest{
			EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
			FieldMask: pbtypes.FieldMask{
				Paths: []string{"ids", "version_ids", "formatters"},
			},
		}, creds)
		a.So(errors.IsNotFound(err), should.BeTrue)
	}
}

func TestDeviceRegistryGet(t *testing.T) {
	type getByIDCallKey struct{}

	registeredDevice = &ttnpb.EndDevice{
		EndDeviceIdentifiers: ttnpb.EndDeviceIdentifiers{
			ApplicationIdentifiers: ttnpb.ApplicationIdentifiers{
				ApplicationID: ApplicationID,
			},
			DeviceID: DeviceID,
			JoinEUI:  eui64Ptr(types.EUI64{0x42, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
			DevEUI:   eui64Ptr(types.EUI64{0x42, 0x42, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
		},
		VersionIDs: &ttnpb.EndDeviceVersionIdentifiers{
			BrandID:         "thethingsproducts",
			ModelID:         "thethingsnode",
			HardwareVersion: "1.0",
			FirmwareVersion: "1.1",
		},
		Formatters: &ttnpb.MessagePayloadFormatters{
			UpFormatter:   ttnpb.PayloadFormatter_FORMATTER_REPOSITORY,
			DownFormatter: ttnpb.PayloadFormatter_FORMATTER_REPOSITORY,
		},
	}

	for _, tc := range []struct {
		Name             string
		ContextFunc      func(context.Context) context.Context
		GetFunc          func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string) (*ttnpb.EndDevice, error)
		Request          *ttnpb.GetEndDeviceRequest
		Device           *ttnpb.EndDevice
		ErrorAssertion   func(*testing.T, error) bool
		ContextAssertion func(context.Context) bool
	}{
		{
			Name: "Unauthorized: no credentials",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, registeredDevice.ApplicationIdentifiers): {
							Rights: []ttnpb.Right{},
						},
					},
				})
			},
			GetFunc: func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string) (*ttnpb.EndDevice, error) {
				test.MustTFromContext(ctx).Errorf("GetFunc must not be called")
				return nil, errors.New("GetFunc must not be called")
			},
			Request: &ttnpb.GetEndDeviceRequest{
				EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
				FieldMask: pbtypes.FieldMask{
					Paths: []string{
						"test",
					},
				},
			},
			ErrorAssertion: func(t *testing.T, err error) bool {
				if !assertions.New(t).So(errors.IsPermissionDenied(err), should.BeTrue) {
					t.Errorf("Received error: %s", err)
					return false
				}
				return true
			},
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, getByIDCallKey{}), should.Equal, 0)
			},
		},

		{
			Name: "Unauthorized: wrong application",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, ttnpb.ApplicationIdentifiers{ApplicationID: "other-app"}): {
							Rights: []ttnpb.Right{
								ttnpb.RIGHT_APPLICATION_DEVICES_READ,
							},
						},
					},
				})
			},
			GetFunc: func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string) (*ttnpb.EndDevice, error) {
				test.MustTFromContext(ctx).Errorf("GetFunc must not be called")
				return nil, errors.New("GetFunc must not be called")
			},
			Request: &ttnpb.GetEndDeviceRequest{
				EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
				FieldMask: pbtypes.FieldMask{
					Paths: []string{
						"test",
					},
				},
			},
			ErrorAssertion: func(t *testing.T, err error) bool {
				if !assertions.New(t).So(errors.IsPermissionDenied(err), should.BeTrue) {
					t.Errorf("Received error: %s", err)
					return false
				}
				return true
			},
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, getByIDCallKey{}), should.Equal, 0)
			},
		},

		{
			Name: "Not found",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, registeredDevice.ApplicationIdentifiers): {
							Rights: []ttnpb.Right{
								ttnpb.RIGHT_APPLICATION_DEVICES_READ,
							},
						},
					},
				})
			},
			GetFunc: func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string) (*ttnpb.EndDevice, error) {
				defer test.MustIncrementContextCounter(ctx, getByIDCallKey{}, 1)
				a := assertions.New(test.MustTFromContext(ctx))
				a.So(ids, should.Resemble, ids)
				a.So(paths, should.HaveSameElementsDeep, []string{"test"})
				return nil, errNotFound
			},
			Request: &ttnpb.GetEndDeviceRequest{
				EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
				FieldMask: pbtypes.FieldMask{
					Paths: []string{
						"test",
					},
				},
			},
			ErrorAssertion: func(t *testing.T, err error) bool {
				if !assertions.New(t).So(errors.IsNotFound(err), should.BeTrue) {
					t.Errorf("Received error: %s", err)
					return false
				}
				return true
			},
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, getByIDCallKey{}), should.Equal, 1)
			},
		},

		{
			Name: "Valid request",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, registeredDevice.ApplicationIdentifiers): {
							Rights: []ttnpb.Right{
								ttnpb.RIGHT_APPLICATION_DEVICES_READ,
							},
						},
					},
				})
			},
			GetFunc: func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string) (*ttnpb.EndDevice, error) {
				a := assertions.New(test.MustTFromContext(ctx))
				a.So(ids, should.Resemble, registeredDevice.EndDeviceIdentifiers)
				a.So(paths, should.HaveSameElementsDeep, []string{
					"test",
				})
				return registeredDevice, nil
			},
			Request: &ttnpb.GetEndDeviceRequest{
				EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
				FieldMask: pbtypes.FieldMask{
					Paths: []string{
						"test",
					},
				},
			},
			Device: registeredDevice,
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, getByIDCallKey{}), should.Equal, 1)
			},
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			a := assertions.New(t)

			as := test.Must(New(component.MustNew(test.GetLogger(t), &component.Config{}),
				&Config{
					LinkMode: "explicit",
					Devices: &MockDeviceRegistry{
						GetFunc: tc.GetFunc,
					},
				})).(*ApplicationServer)

			as.AddContextFiller(tc.ContextFunc)
			as.AddContextFiller(func(ctx context.Context) context.Context {
				return test.ContextWithCounter(ctx, getByIDCallKey{})
			})
			as.AddContextFiller(func(ctx context.Context) context.Context {
				ctx, cancel := context.WithDeadline(ctx, time.Now().Add(Timeout))
				_ = cancel
				return ctx
			})
			as.AddContextFiller(func(ctx context.Context) context.Context {
				return test.ContextWithT(ctx, t)
			})
			test.Must(nil, as.Start())
			defer as.Close()

			req := deepcopy.Copy(tc.Request).(*ttnpb.GetEndDeviceRequest)

			dev, err := ttnpb.NewAsEndDeviceRegistryClient(as.LoopbackConn()).Get(test.Context(), req)
			if tc.ErrorAssertion != nil {
				a.So(tc.ErrorAssertion(t, err), should.BeTrue)
				a.So(dev, should.BeNil)
			} else {
				a.So(err, should.BeNil)
				a.So(dev, should.Resemble, tc.Device)
			}
			a.So(req, should.Resemble, tc.Request)
		})
	}
}

func TestDeviceRegistrySet(t *testing.T) {
	type setByIDCallKey struct{}

	registeredDevice = &ttnpb.EndDevice{
		EndDeviceIdentifiers: ttnpb.EndDeviceIdentifiers{
			ApplicationIdentifiers: ttnpb.ApplicationIdentifiers{
				ApplicationID: ApplicationID,
			},
			DeviceID: DeviceID,
			JoinEUI:  eui64Ptr(types.EUI64{0x42, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
			DevEUI:   eui64Ptr(types.EUI64{0x42, 0x42, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
		},
		VersionIDs: &ttnpb.EndDeviceVersionIdentifiers{
			BrandID:         "thethingsproducts",
			ModelID:         "thethingsnode",
			HardwareVersion: "1.0",
			FirmwareVersion: "1.1",
		},
		Formatters: &ttnpb.MessagePayloadFormatters{
			UpFormatter:   ttnpb.PayloadFormatter_FORMATTER_REPOSITORY,
			DownFormatter: ttnpb.PayloadFormatter_FORMATTER_REPOSITORY,
		},
	}

	for _, tc := range []struct {
		Name             string
		ContextFunc      func(context.Context) context.Context
		SetFunc          func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error)
		Request          *ttnpb.SetEndDeviceRequest
		Device           *ttnpb.EndDevice
		ErrorAssertion   func(*testing.T, error) bool
		ContextAssertion func(context.Context) bool
	}{
		{
			Name: "Unauthorized: no credentials",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, registeredDevice.ApplicationIdentifiers): {
							Rights: []ttnpb.Right{},
						},
					},
				})
			},
			SetFunc: func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error) {
				test.MustTFromContext(ctx).Errorf("SetFunc must not be called")
				return nil, errors.New("SetFunc must not be called")
			},
			Request: &ttnpb.SetEndDeviceRequest{
				Device: ttnpb.EndDevice{
					EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
					FrequencyPlanID:      "EU_863_870",
				},
				FieldMask: pbtypes.FieldMask{
					Paths: []string{"frequency_plan_id"},
				},
			},
			ErrorAssertion: func(t *testing.T, err error) bool {
				if !assertions.New(t).So(errors.IsPermissionDenied(err), should.BeTrue) {
					t.Errorf("Received error: %s", err)
					return false
				}
				return true
			},
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, setByIDCallKey{}), should.Equal, 0)
			},
		},

		{
			Name: "Unauthorized: wrong application",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, ttnpb.ApplicationIdentifiers{ApplicationID: "other-app"}): {
							Rights: []ttnpb.Right{},
						},
					},
				})
			},
			SetFunc: func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error) {
				test.MustTFromContext(ctx).Errorf("SetFunc must not be called")
				return nil, errors.New("SetFunc must not be called")
			},
			Request: &ttnpb.SetEndDeviceRequest{
				Device: ttnpb.EndDevice{
					EndDeviceIdentifiers: registeredDevice.EndDeviceIdentifiers,
					FrequencyPlanID:      "EU_863_870",
				},
				FieldMask: pbtypes.FieldMask{
					Paths: []string{"frequency_plan_id"},
				},
			},
			ErrorAssertion: func(t *testing.T, err error) bool {
				if !assertions.New(t).So(errors.IsPermissionDenied(err), should.BeTrue) {
					t.Errorf("Received error: %s", err)
					return false
				}
				return true
			},
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, setByIDCallKey{}), should.Equal, 0)
			},
		},

		{
			Name: "Create",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, registeredDevice.ApplicationIdentifiers): {
							Rights: []ttnpb.Right{
								ttnpb.RIGHT_APPLICATION_DEVICES_WRITE,
							},
						},
					},
				})
			},
			SetFunc: func(ctx context.Context, deviceIds ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error) {
				a := assertions.New(test.MustTFromContext(ctx))
				defer test.MustIncrementContextCounter(ctx, setByIDCallKey{}, 1)
				a.So(deviceIds, should.Resemble, registeredDevice.EndDeviceIdentifiers)
				a.So(paths, should.Contain, "frequency_plan_id")
				dev, _, err := f(deepcopy.Copy(registeredDevice).(*ttnpb.EndDevice))
				test.MustTFromContext(ctx).Log(dev)
				return dev, err
			},
			Request: &ttnpb.SetEndDeviceRequest{
				Device: *registeredDevice,
				FieldMask: pbtypes.FieldMask{
					Paths: []string{"frequency_plan_id"},
				},
			},
			Device: registeredDevice,
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, setByIDCallKey{}), should.Equal, 1)
			},
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			a := assertions.New(t)

			as := test.Must(New(component.MustNew(test.GetLogger(t), &component.Config{}),
				&Config{
					LinkMode: "explicit",
					Devices: &MockDeviceRegistry{
						SetFunc: tc.SetFunc,
					},
				})).(*ApplicationServer)

			as.AddContextFiller(tc.ContextFunc)
			as.AddContextFiller(func(ctx context.Context) context.Context {
				return test.ContextWithCounter(ctx, setByIDCallKey{})
			})
			as.AddContextFiller(func(ctx context.Context) context.Context {
				ctx, cancel := context.WithDeadline(ctx, time.Now().Add(Timeout))
				_ = cancel
				return ctx
			})
			as.AddContextFiller(func(ctx context.Context) context.Context {
				return test.ContextWithT(ctx, t)
			})
			test.Must(nil, as.Start())
			defer as.Close()

			req := deepcopy.Copy(tc.Request).(*ttnpb.SetEndDeviceRequest)

			dev, err := ttnpb.NewAsEndDeviceRegistryClient(as.LoopbackConn()).Set(test.Context(), req)
			if tc.ErrorAssertion != nil {
				a.So(tc.ErrorAssertion(t, err), should.BeTrue)
				a.So(dev, should.BeNil)
			} else {
				a.So(err, should.BeNil)
				a.So(dev, should.Resemble, tc.Device)
			}
			a.So(req, should.Resemble, tc.Request)
		})
	}
}

func TestDeviceRegistryDelete(t *testing.T) {
	type setByIDCallKey struct{}

	registeredDevice = &ttnpb.EndDevice{
		EndDeviceIdentifiers: ttnpb.EndDeviceIdentifiers{
			ApplicationIdentifiers: ttnpb.ApplicationIdentifiers{
				ApplicationID: ApplicationID,
			},
			DeviceID: DeviceID,
			JoinEUI:  eui64Ptr(types.EUI64{0x42, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
			DevEUI:   eui64Ptr(types.EUI64{0x42, 0x42, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
		},
		VersionIDs: &ttnpb.EndDeviceVersionIdentifiers{
			BrandID:         "thethingsproducts",
			ModelID:         "thethingsnode",
			HardwareVersion: "1.0",
			FirmwareVersion: "1.1",
		},
		Formatters: &ttnpb.MessagePayloadFormatters{
			UpFormatter:   ttnpb.PayloadFormatter_FORMATTER_REPOSITORY,
			DownFormatter: ttnpb.PayloadFormatter_FORMATTER_REPOSITORY,
		},
	}

	for _, tc := range []struct {
		Name             string
		ContextFunc      func(context.Context) context.Context
		SetFunc          func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error)
		Request          *ttnpb.EndDeviceIdentifiers
		ErrorAssertion   func(*testing.T, error) bool
		ContextAssertion func(context.Context) bool
	}{
		{
			Name: "Unauthorized: no credentials",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, registeredDevice.ApplicationIdentifiers): {
							Rights: []ttnpb.Right{},
						},
					},
				})
			},
			SetFunc: func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error) {
				test.MustTFromContext(ctx).Errorf("SetFunc must not be called")
				return nil, errors.New("SetFunc must not be called")
			},
			Request: deepcopy.Copy(&registeredDevice.EndDeviceIdentifiers).(*ttnpb.EndDeviceIdentifiers),
			ErrorAssertion: func(t *testing.T, err error) bool {
				if !assertions.New(t).So(errors.IsPermissionDenied(err), should.BeTrue) {
					t.Errorf("Received error: %s", err)
					return false
				}
				return true
			},
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, setByIDCallKey{}), should.Equal, 0)
			},
		},

		{
			Name: "Unauthorized: wrong application",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, ttnpb.ApplicationIdentifiers{ApplicationID: "other-app"}): {
							Rights: []ttnpb.Right{},
						},
					},
				})
			},
			SetFunc: func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error) {
				test.MustTFromContext(ctx).Errorf("SetFunc must not be called")
				return nil, errors.New("SetFunc must not be called")
			},
			Request: deepcopy.Copy(&registeredDevice.EndDeviceIdentifiers).(*ttnpb.EndDeviceIdentifiers),
			ErrorAssertion: func(t *testing.T, err error) bool {
				if !assertions.New(t).So(errors.IsPermissionDenied(err), should.BeTrue) {
					t.Errorf("Received error: %s", err)
					return false
				}
				return true
			},
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, setByIDCallKey{}), should.Equal, 0)
			},
		},

		{
			Name: "Non-existing device",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, registeredDevice.ApplicationIdentifiers): {
							Rights: []ttnpb.Right{
								ttnpb.RIGHT_APPLICATION_DEVICES_WRITE,
							},
						},
					},
				})
			},
			SetFunc: func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error) {
				defer test.MustIncrementContextCounter(ctx, setByIDCallKey{}, 1)
				t := test.MustTFromContext(ctx)
				a := assertions.New(t)
				a.So(ids, should.Resemble, registeredDevice.EndDeviceIdentifiers)

				dev, sets, err := f(nil)
				a.So(err, should.BeNil)
				a.So(sets, should.BeNil)
				a.So(dev, should.BeNil)
				return nil, nil
			},
			Request: deepcopy.Copy(&registeredDevice.EndDeviceIdentifiers).(*ttnpb.EndDeviceIdentifiers),
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, setByIDCallKey{}), should.Equal, 1)
			},
		},

		{
			Name: "Existing device",
			ContextFunc: func(ctx context.Context) context.Context {
				return rights.NewContext(ctx, rights.Rights{
					ApplicationRights: map[string]*ttnpb.Rights{
						unique.ID(ctx, registeredDevice.ApplicationIdentifiers): {
							Rights: []ttnpb.Right{
								ttnpb.RIGHT_APPLICATION_DEVICES_WRITE,
							},
						},
					},
				})
			},
			SetFunc: func(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error) {
				defer test.MustIncrementContextCounter(ctx, setByIDCallKey{}, 1)
				t := test.MustTFromContext(ctx)
				a := assertions.New(t)
				a.So(ids, should.Resemble, registeredDevice.EndDeviceIdentifiers)

				dev, sets, err := f(registeredDevice)
				a.So(err, should.BeNil)
				a.So(sets, should.BeNil)
				a.So(dev, should.BeNil)
				return nil, nil
			},
			Request: deepcopy.Copy(&registeredDevice.EndDeviceIdentifiers).(*ttnpb.EndDeviceIdentifiers),
			ContextAssertion: func(ctx context.Context) bool {
				a := assertions.New(test.MustTFromContext(ctx))
				return a.So(test.MustCounterFromContext(ctx, setByIDCallKey{}), should.Equal, 1)
			},
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			a := assertions.New(t)

			as := test.Must(New(component.MustNew(test.GetLogger(t), &component.Config{}),
				&Config{
					LinkMode: "explicit",
					Devices: &MockDeviceRegistry{
						SetFunc: tc.SetFunc,
					},
				})).(*ApplicationServer)

			as.AddContextFiller(tc.ContextFunc)
			as.AddContextFiller(func(ctx context.Context) context.Context {
				return test.ContextWithCounter(ctx, setByIDCallKey{})
			})
			as.AddContextFiller(func(ctx context.Context) context.Context {
				ctx, cancel := context.WithDeadline(ctx, time.Now().Add(Timeout))
				_ = cancel
				return ctx
			})
			as.AddContextFiller(func(ctx context.Context) context.Context {
				return test.ContextWithT(ctx, t)
			})
			test.Must(nil, as.Start())
			defer as.Close()

			req := deepcopy.Copy(tc.Request).(*ttnpb.EndDeviceIdentifiers)

			res, err := ttnpb.NewAsEndDeviceRegistryClient(as.LoopbackConn()).Delete(test.Context(), req)
			if tc.ErrorAssertion != nil {
				a.So(tc.ErrorAssertion(t, err), should.BeTrue)
				a.So(res, should.BeNil)
			} else {
				a.So(err, should.BeNil)
				a.So(res, should.Resemble, ttnpb.Empty)
			}
			a.So(req, should.Resemble, tc.Request)
		})
	}
}
