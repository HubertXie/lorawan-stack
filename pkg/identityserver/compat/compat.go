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
	"context"

	"github.com/labstack/echo"
	"go.thethings.network/lorawan-stack/pkg/frequencyplans"
	"go.thethings.network/lorawan-stack/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/web"
)

// Server implements the legacy endpoints (the gateway information endpoint and the
// frequency plan information endpoint) used by The Things Gateway.
type Server struct {
	fpStore  *frequencyplans.Store
	gtwStore store.GatewayStore

	ctx context.Context
}

const compatAPIPrefix = "/api/v2"

// RegisterRoutes implements the web.Registerer interface.
func (s *Server) RegisterRoutes(srv *web.Server) {
	group := srv.Group(compatAPIPrefix)
	gatewayMiddleware := []echo.MiddlewareFunc{
		s.validateAndFillGatewayIDs(),
		s.requireGatewayRights(ttnpb.RIGHT_GATEWAY_INFO),
	}
	group.GET("/gateways/:gateway_id", func(c echo.Context) error {
		return s.handleGatewayInfo(c)
	}, gatewayMiddleware...)
	group.GET("/frequency-plans/:frequency_plan_id", func(c echo.Context) error {
		return s.handleFreqPlanInfo(c)
	})
}

// NewServer creats a new The Things Gateway compatibility server.
func NewServer(ctx context.Context, fpStore *frequencyplans.Store, gtwStore store.GatewayStore) *Server {
	return &Server{
		ctx:      ctx,
		fpStore:  fpStore,
		gtwStore: gtwStore,
	}
}
