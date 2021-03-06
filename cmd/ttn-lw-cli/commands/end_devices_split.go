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

package commands

import (
	"context"

	"github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/pkg/errors"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

func selectPath(path string, allowedPaths []string) bool {
	for _, allowedPath := range allowedPaths {
		if path == allowedPath {
			return true
		}
	}
	return false
}

var (
	getEndDeviceFromIS = ttnpb.AllowedFieldMaskPathsForRPC["/ttn.lorawan.v3.EndDeviceRegistry/Get"]
	getEndDeviceFromNS = ttnpb.AllowedFieldMaskPathsForRPC["/ttn.lorawan.v3.NsEndDeviceRegistry/Get"]
	getEndDeviceFromAS = ttnpb.AllowedFieldMaskPathsForRPC["/ttn.lorawan.v3.AsEndDeviceRegistry/Get"]
	getEndDeviceFromJS = ttnpb.AllowedFieldMaskPathsForRPC["/ttn.lorawan.v3.JsEndDeviceRegistry/Get"]
	setEndDeviceToIS   = ttnpb.AllowedFieldMaskPathsForRPC["/ttn.lorawan.v3.EndDeviceRegistry/Update"]
	setEndDeviceToNS   = ttnpb.AllowedFieldMaskPathsForRPC["/ttn.lorawan.v3.NsEndDeviceRegistry/Set"]
	setEndDeviceToAS   = ttnpb.AllowedFieldMaskPathsForRPC["/ttn.lorawan.v3.AsEndDeviceRegistry/Set"]
	setEndDeviceToJS   = ttnpb.AllowedFieldMaskPathsForRPC["/ttn.lorawan.v3.JsEndDeviceRegistry/Set"]
)

func selectPaths(paths []string, allowedPaths []string) []string {
	selectedPaths := make([]string, 0, len(paths))
	for _, path := range paths {
		if selectPath(path, allowedPaths) {
			selectedPaths = append(selectedPaths, path)
			continue
		}
	}
	return selectedPaths
}

func splitEndDeviceGetPaths(paths ...string) (is, ns, as, js []string) {
	is = selectPaths(paths, getEndDeviceFromIS)
	ns = selectPaths(paths, getEndDeviceFromNS)
	as = selectPaths(paths, getEndDeviceFromAS)
	js = selectPaths(paths, getEndDeviceFromJS)
	return
}

func splitEndDeviceSetPaths(supportsJoin bool, paths ...string) (is, ns, as, js []string) {
	is = selectPaths(paths, setEndDeviceToIS)
	ns = selectPaths(paths, setEndDeviceToNS)
	as = selectPaths(paths, setEndDeviceToAS)
	if supportsJoin {
		js = selectPaths(paths, setEndDeviceToJS)
	}
	return
}

func getEndDevice(ids ttnpb.EndDeviceIdentifiers, nsPaths, asPaths, jsPaths []string, continueOnError bool) (*ttnpb.EndDevice, error) {
	var res ttnpb.EndDevice

	if len(jsPaths) > 0 {
		js, err := api.Dial(ctx, config.JoinServerAddress)
		if err != nil {
			if !continueOnError {
				return nil, err
			}
			logger.WithError(err).Error("Could not connect to Join Server")
		} else {
			jsRes, err := ttnpb.NewJsEndDeviceRegistryClient(js).Get(ctx, &ttnpb.GetEndDeviceRequest{
				EndDeviceIdentifiers: ids,
				FieldMask:            types.FieldMask{Paths: jsPaths},
			})
			if err != nil {
				if !continueOnError {
					return nil, err
				}
				logger.WithError(err).Error("Could not get end device from Join Server")
			} else {
				res.SetFields(jsRes, jsPaths...)
			}
		}
	}

	if len(asPaths) > 0 {
		as, err := api.Dial(ctx, config.ApplicationServerAddress)
		if err != nil {
			if !continueOnError {
				return nil, err
			}
			logger.WithError(err).Error("Could not connect to Application Server")
		} else {
			asRes, err := ttnpb.NewAsEndDeviceRegistryClient(as).Get(ctx, &ttnpb.GetEndDeviceRequest{
				EndDeviceIdentifiers: ids,
				FieldMask:            types.FieldMask{Paths: asPaths},
			})
			if err != nil {
				if !continueOnError {
					return nil, err
				}
				logger.WithError(err).Error("Could not get end device from Application Server")
			} else {
				res.SetFields(asRes, asPaths...)
			}
		}
	}

	if len(nsPaths) > 0 {
		ns, err := api.Dial(ctx, config.NetworkServerAddress)
		if err != nil {
			if !continueOnError {
				return nil, err
			}
			logger.WithError(err).Error("Could not connect to Network Server")
		} else {
			nsRes, err := ttnpb.NewNsEndDeviceRegistryClient(ns).Get(ctx, &ttnpb.GetEndDeviceRequest{
				EndDeviceIdentifiers: ids,
				FieldMask:            types.FieldMask{Paths: nsPaths},
			})
			if err != nil {
				if !continueOnError {
					return nil, err
				}
				logger.WithError(err).Error("Could not get end device from Network Server")
			} else {
				res.SetFields(nsRes, nsPaths...)
			}
		}
	}

	return &res, nil
}

func setEndDevice(device *ttnpb.EndDevice, isPaths, nsPaths, asPaths, jsPaths []string, isCreate bool) (*ttnpb.EndDevice, error) {
	var res ttnpb.EndDevice

	if len(isPaths) > 0 || isCreate {
		is, err := api.Dial(ctx, config.IdentityServerAddress)
		if err != nil {
			return nil, err
		}
		var isDevice ttnpb.EndDevice
		isDevice.SetFields(device, append(isPaths, "ids")...)
		isRes, err := ttnpb.NewEndDeviceRegistryClient(is).Update(ctx, &ttnpb.UpdateEndDeviceRequest{
			EndDevice: isDevice,
			FieldMask: types.FieldMask{Paths: isPaths},
		})
		if err != nil {
			return nil, err
		}
		res.SetFields(isRes, isPaths...)
	}

	if len(jsPaths) > 0 {
		js, err := api.Dial(ctx, config.JoinServerAddress)
		if err != nil {
			return nil, err
		}
		var jsDevice ttnpb.EndDevice
		jsDevice.SetFields(device, append(jsPaths, "ids")...)
		jsRes, err := ttnpb.NewJsEndDeviceRegistryClient(js).Set(ctx, &ttnpb.SetEndDeviceRequest{
			EndDevice: jsDevice,
			FieldMask: types.FieldMask{Paths: jsPaths},
		})
		if err != nil {
			return nil, err
		}
		res.SetFields(jsRes, jsPaths...)
	}

	if len(nsPaths) > 0 || isCreate {
		ns, err := api.Dial(ctx, config.NetworkServerAddress)
		if err != nil {
			return nil, err
		}
		var nsDevice ttnpb.EndDevice
		nsDevice.SetFields(device, append(nsPaths, "ids")...)
		nsRes, err := ttnpb.NewNsEndDeviceRegistryClient(ns).Set(ctx, &ttnpb.SetEndDeviceRequest{
			EndDevice: nsDevice,
			FieldMask: types.FieldMask{Paths: nsPaths},
		})
		if err != nil {
			return nil, err
		}
		res.SetFields(nsRes, nsPaths...)
	}

	if len(asPaths) > 0 || isCreate {
		as, err := api.Dial(ctx, config.ApplicationServerAddress)
		if err != nil {
			return nil, err
		}
		var asDevice ttnpb.EndDevice
		asDevice.SetFields(device, append(asPaths, "ids")...)
		asRes, err := ttnpb.NewAsEndDeviceRegistryClient(as).Set(ctx, &ttnpb.SetEndDeviceRequest{
			EndDevice: asDevice,
			FieldMask: types.FieldMask{Paths: asPaths},
		})
		if err != nil {
			return nil, err
		}
		res.SetFields(asRes, asPaths...)
	}

	return &res, ctx.Err()
}

func deleteEndDevice(ctx context.Context, devID *ttnpb.EndDeviceIdentifiers) error {
	as, err := api.Dial(ctx, config.ApplicationServerAddress)
	if err != nil {
		return err
	}
	_, err = ttnpb.NewAsEndDeviceRegistryClient(as).Delete(ctx, devID)
	if errors.IsNotFound(err) {
		logger.WithError(err).Error("Could not delete end device from Application Server")
	} else if err != nil {
		return err
	}

	ns, err := api.Dial(ctx, config.NetworkServerAddress)
	if err != nil {
		return err
	}
	_, err = ttnpb.NewNsEndDeviceRegistryClient(ns).Delete(ctx, devID)
	if errors.IsNotFound(err) {
		logger.WithError(err).Error("Could not delete end device from Network Server")
	} else if err != nil {
		return err
	}

	if devID.JoinEUI != nil && devID.DevEUI != nil {
		js, err := api.Dial(ctx, config.JoinServerAddress)
		if err != nil {
			return err
		}
		_, err = ttnpb.NewJsEndDeviceRegistryClient(js).Delete(ctx, devID)
		if errors.IsNotFound(err) {
			logger.WithError(err).Error("Could not delete end device from Join Server")
		} else if err != nil {
			return err
		}
	}

	is, err := api.Dial(ctx, config.IdentityServerAddress)
	if err != nil {
		return err
	}
	_, err = ttnpb.NewEndDeviceRegistryClient(is).Delete(ctx, devID)
	if err != nil {
		return err
	}

	return nil
}
