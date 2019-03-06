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
	"testing"

	"github.com/smartystreets/assertions"
)

func TestAdaptAuthorization(t *testing.T) {
	for _, tt := range []struct {
		Name          string
		Authorization string
		Assert        func(*assertions.Assertion, string, error)
	}{
		{
			Name:          "Empty Authorization",
			Authorization: "",
			Assert: func(a *assertions.Assertion, auth string, err error) {
				a.So(auth, assertions.ShouldEqual, "")
				a.So(err, assertions.ShouldBeNil)
			},
		},
		{
			Name:          "Key formatted authorization",
			Authorization: "Key asd",
			Assert: func(a *assertions.Assertion, auth string, err error) {
				a.So(auth, assertions.ShouldEqual, "Bearer asd")
				a.So(err, assertions.ShouldBeNil)
			},
		},
		{
			Name:          "Bearer formatted authorization",
			Authorization: "Bearer efg",
			Assert: func(a *assertions.Assertion, auth string, err error) {
				a.So(auth, assertions.ShouldEqual, "Bearer efg")
				a.So(err, assertions.ShouldBeNil)
			},
		},
		{
			Name:          "Invalid authorization",
			Authorization: "InvalidKeyFormat",
			Assert: func(a *assertions.Assertion, auth string, err error) {
				a.So(auth, assertions.ShouldEqual, "")
				a.So(err, assertions.ShouldNotBeNil)
			},
		},
		{
			Name:          "Esoteric authorization",
			Authorization: "APIKey asd",
			Assert: func(a *assertions.Assertion, auth string, err error) {
				a.So(auth, assertions.ShouldEqual, "Bearer asd")
				a.So(err, assertions.ShouldBeNil)
			},
		},
	} {
		t.Run(tt.Name, func(t *testing.T) {
			a := assertions.New(t)
			auth, err := adaptAuthorization(tt.Authorization)
			tt.Assert(a, auth, err)
		})
	}
}
