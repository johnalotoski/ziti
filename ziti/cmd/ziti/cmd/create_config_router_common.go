/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package cmd

import (
	_ "embed"
	cmdhelper "github.com/openziti/ziti/ziti/cmd/ziti/cmd/helpers"
	"os"
)

func SetZitiRouterIdentity(r *RouterTemplateValues, routerName string) {
	SetZitiRouterIdentityCert(r, routerName)
	SetZitiRouterIdentityServerCert(r, routerName)
	SetZitiRouterIdentityKey(r, routerName)
	SetZitiRouterIdentityCA(r, routerName)
}
func SetZitiRouterIdentityCert(r *RouterTemplateValues, routerName string) {
	val := os.Getenv("ZITI_ROUTER_IDENTITY_CERT")
	if val == "" {
		val = workingDir + "/" + routerName + ".cert" //default
	}
	r.IdentityCert = cmdhelper.NormalizePath(val)
}
func SetZitiRouterIdentityServerCert(r *RouterTemplateValues, routerName string) {
	val := os.Getenv("ZITI_ROUTER_IDENTITY_SERVER_CERT")
	if val == "" {
		val = workingDir + "/" + routerName + ".server.chain.cert" //default
	}
	r.IdentityServerCert = cmdhelper.NormalizePath(val)
}
func SetZitiRouterIdentityKey(r *RouterTemplateValues, routerName string) {
	val := os.Getenv("ZITI_ROUTER_IDENTITY_KEY")
	if val == "" {
		val = workingDir + "/" + routerName + ".key" //default
	}
	r.IdentityKey = cmdhelper.NormalizePath(val)
}
func SetZitiRouterIdentityCA(r *RouterTemplateValues, routerName string) {
	val := os.Getenv("ZITI_ROUTER_IDENTITY_CA")
	if val == "" {
		val = workingDir + "/" + routerName + ".cas" //default
	}
	r.IdentityCA = cmdhelper.NormalizePath(val)
}
