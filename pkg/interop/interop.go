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

package interop

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	echomiddleware "github.com/labstack/echo/middleware"
	"go.thethings.network/lorawan-stack/pkg/config"
	"go.thethings.network/lorawan-stack/pkg/log"
	"go.thethings.network/lorawan-stack/pkg/web"
	"go.thethings.network/lorawan-stack/pkg/web/middleware"
)

// Registerer allows components to register their interop services to the web server.
type Registerer interface {
	RegisterInterop(s *Server)
}

// JoinServer represents a Join Server.
type JoinServer interface {
	JoinRequest(req *JoinReq) (*JoinAns, error)
}

// HomeNetworkServer represents a Home Network Server.
type HomeNetworkServer interface {
}

// ServingNetworkServer represents a Serving Network Server.
type ServingNetworkServer interface {
}

// ForwardingNetworkServer represents a Forwarding Network Server.
type ForwardingNetworkServer interface {
}

// ApplicationServer represents an Application Server.
type ApplicationServer interface {
}

// Server is the server.
type Server struct {
	rootGroup *echo.Group
	server    *echo.Echo
	config    config.Interop

	js  JoinServer
	hNs HomeNetworkServer
	sNs ServingNetworkServer
	fNs ForwardingNetworkServer
	as  ApplicationServer
}

// New builds a new server.
func New(ctx context.Context, config config.Interop) (*Server, error) {
	logger := log.FromContext(ctx).WithField("namespace", "interop")

	server := echo.New()

	server.Logger = web.NewNoopLogger()
	server.HTTPErrorHandler = ErrorHandler

	server.Use(
		middleware.ID("interop"),
		echomiddleware.BodyLimit("16M"),
		echomiddleware.Secure(),
		middleware.Recover(),
	)

	s := &Server{
		rootGroup: server.Group(
			"",
			middleware.Log(logger),
			middleware.Normalize(middleware.RedirectPermanent),
		),
		config: config,
		server: server,
	}

	// In 1.0, NS, JS and AS receive messages on the root path.
	// In 1.1, only JS and AS receive messages on the root path. Since NS can play various roles (hNS, sNS and fNS), their
	// group is created on registration of the handler.
	s.rootGroup.POST("/", s.handleRequest)

	return s, nil
}

// ServeHTTP implements http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.server.ServeHTTP(w, r)
}

// RegisterJs registers the Join Server for AS-JS, hNS-JS and vNS-JS messages.
func (s *Server) RegisterJs(js JoinServer) {
	s.js = js
}

// RegisterHNs registers the Home Network Server for AS-hNS, JS-hNS and sNS-hNS messages.
func (s *Server) RegisterHNs(hNs HomeNetworkServer) {
	s.hNs = hNs
	s.rootGroup.POST("/hns", s.handleNsRequest)
}

// RegisterSNs registers the Serving Network Server for hNS-sNS, fNS-sNS and JS-vNS messages.
func (s *Server) RegisterSNs(sNs ServingNetworkServer) {
	s.sNs = sNs
	s.rootGroup.POST("/sns", s.handleNsRequest)
}

// RegisterFNs registers the Forwarding Network Server for sNS-fNS and JS-vNS messages.
func (s *Server) RegisterFNs(fNs ForwardingNetworkServer) {
	s.fNs = fNs
	s.rootGroup.POST("/fns", s.handleNsRequest)
}

// RegisterAs registers the Application Server for JS-AS messages.
func (s *Server) RegisterAs(as ApplicationServer) {
	s.as = as
}

func (s *Server) handleRequest(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	var header MessageHeader
	if err := json.Unmarshal(body, &header); err != nil {
		return err
	}
	return nil
}

func (s *Server) handleNsRequest(c echo.Context) error {
	// TODO: Implement LoRaWAN roaming (https://github.com/TheThingsNetwork/lorawan-stack/issues/230)
	c.NoContent(http.StatusNotFound)
	return nil
}
