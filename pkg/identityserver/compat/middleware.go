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
	"strings"

	"github.com/labstack/echo"
	"go.thethings.network/lorawan-stack/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"google.golang.org/grpc/metadata"
)

// validateAndFillGatewayIDs checks if the request contains a valid gateway ID.
func (s *Server) validateAndFillGatewayIDs() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			gatewayIDs := ttnpb.GatewayIdentifiers{
				GatewayID: c.Param(gatewayIDKey),
			}
			if err := gatewayIDs.ValidateContext(s.ctx); err != nil {
				return err
			}
			c.Set(gatewayIDKey, gatewayIDs)

			return next(c)
		}
	}
}

// requestGatewayRights checks if the request contains a valid API key which has
// the required rights over the gateway.
func (s *Server) requireGatewayRights(required ...ttnpb.Right) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := s.ctx
			gatewayIDs := c.Get(gatewayIDKey).(ttnpb.GatewayIdentifiers)
			auth := c.Request().Header.Get(echo.HeaderAuthorization)
			md := metadata.New(map[string]string{
				"id":            gatewayIDs.GatewayID,
				"authorization": adaptAuthorization(auth),
			})
			if ctxMd, ok := metadata.FromIncomingContext(ctx); ok {
				md = metadata.Join(ctxMd, md)
			}
			ctx = metadata.NewIncomingContext(ctx, md)
			if err := rights.RequireGateway(ctx, gatewayIDs, required...); err != nil {
				return err
			}
			return next(c)
		}
	}
}

const (
	oldAuthKey = "Key"
	newAuthKey = "Bearer"
)

// adaptAuthorization returns converts the Authorization header value to the
// "Key" format to the "Bearer" format, if needed.
func adaptAuthorization(originalAuth string) string {
	if !strings.HasPrefix(originalAuth, oldAuthKey) {
		return originalAuth
	}
	return newAuthKey + strings.TrimPrefix(originalAuth, oldAuthKey)
}
