/*
	Copyright NetFoundry Inc.

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

package model

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/openziti/ziti/common/cert"
	"github.com/openziti/ziti/common/eid"
	"github.com/openziti/ziti/controller/config"
	"github.com/openziti/ziti/controller/jwtsigner"
	"github.com/openziti/ziti/controller/persistence"
	"github.com/openziti/ziti/controller/change"
	"github.com/openziti/ziti/controller/network"
	"github.com/openziti/identity"
	"github.com/openziti/metrics"
	"testing"
	"time"
)

var _ Env = &TestContext{}

var _ HostController = &testHostController{}

type testHostController struct {
	closeNotify chan struct{}
	ctx         *persistence.TestContext
}

func (self *testHostController) GetPeerSigners() []*x509.Certificate {
	return nil
}

func (self *testHostController) Identity() identity.Identity {
	return &identity.TokenId{Token: "test"}
}

func (self *testHostController) GetNetwork() *network.Network {
	return self.ctx.GetNetwork()
}

func (self testHostController) Shutdown() {
	close(self.closeNotify)
}

func (self testHostController) GetCloseNotifyChannel() <-chan struct{} {
	return self.closeNotify
}

func (self testHostController) Stop() {
	close(self.closeNotify)
}

func (ctx testHostController) IsRaftEnabled() bool {
	return false
}

type TestContext struct {
	*persistence.TestContext
	managers        *Managers
	config          *config.Config
	metricsRegistry metrics.Registry
	hostController  *testHostController
}

func (ctx *TestContext) JwtSignerKeyFunc(token *jwt.Token) (interface{}, error) {
	tlsCert, _, _ := ctx.GetServerCert()
	return tlsCert.Leaf.PublicKey, nil
}

func (ctx *TestContext) GetServerCert() (*tls.Certificate, string, jwt.SigningMethod) {
	return nil, "", nil
}

func (ctx *TestContext) HandleServiceUpdatedEventForIdentityId(identityId string) {}

func (ctx *TestContext) Generate(string, string, jwt.Claims) (string, error) {
	return "I'm a very legitimate claim", nil
}

func (ctx *TestContext) GetManagers() *Managers {
	return ctx.managers
}

func (ctx *TestContext) GetConfig() *config.Config {
	return ctx.config
}

func (ctx *TestContext) GetJwtSigner() jwtsigner.Signer {
	return ctx
}

func (ctx *TestContext) GetAuthRegistry() AuthRegistry {
	panic("implement me")
}

func (ctx *TestContext) GetEnrollRegistry() EnrollmentRegistry {
	panic("implement me")
}

func (ctx *TestContext) GetApiClientCsrSigner() cert.Signer {
	panic("implement me")
}

func (ctx *TestContext) GetApiServerCsrSigner() cert.Signer {
	panic("implement me")
}

func (ctx *TestContext) GetControlClientCsrSigner() cert.Signer {
	panic("implement me")
}

func (ctx *TestContext) GetHostController() HostController {
	return ctx.hostController
}

func (ctx *TestContext) GetSchemas() Schemas {
	panic("implement me")
}

func (ctx *TestContext) IsEdgeRouterOnline(string) bool {
	panic("implement me")
}

func (ctx *TestContext) GetMetricsRegistry() metrics.Registry {
	return ctx.metricsRegistry
}

func (ctx *TestContext) GetFingerprintGenerator() cert.FingerprintGenerator {
	return nil
}

func NewTestContext(t *testing.T) *TestContext {
	fabricTestContext := persistence.NewTestContext(t)
	context := &TestContext{
		TestContext:     fabricTestContext,
		metricsRegistry: metrics.NewRegistry("test", nil),
		hostController: &testHostController{
			ctx:         fabricTestContext,
			closeNotify: make(chan struct{}),
		},
	}
	return context
}
func (ctx *TestContext) Init() {
	ctx.TestContext.Init()
	ctx.config = &config.Config{
		Enrollment: config.Enrollment{
			EdgeRouter: config.EnrollmentOption{
				Duration: 60 * time.Second,
			},
		},
	}
	ctx.managers = InitEntityManagers(ctx)
}

func (ctx *TestContext) Cleanup() {
	if ctx.hostController != nil {
		ctx.hostController.Stop()
	}
	ctx.TestContext.Cleanup()
}

func (ctx *TestContext) requireNewIdentity(isAdmin bool) *Identity {
	newIdentity := &Identity{
		Name:           eid.New(),
		IsAdmin:        isAdmin,
		IdentityTypeId: persistence.DefaultIdentityType,
	}
	ctx.NoError(ctx.managers.Identity.Create(newIdentity, change.New()))
	return newIdentity
}

func (ctx *TestContext) requireNewService() *Service {
	service := &Service{
		Name: eid.New(),
	}
	ctx.NoError(ctx.managers.EdgeService.Create(service, change.New()))
	return service
}

func (ctx *TestContext) requireNewEdgeRouter() *EdgeRouter {
	edgeRouter := &EdgeRouter{
		Name: eid.New(),
	}
	ctx.NoError(ctx.managers.EdgeRouter.Create(edgeRouter, change.New()))
	return edgeRouter
}

func (ctx *TestContext) requireNewApiSession(identity *Identity) *ApiSession {
	entity := &ApiSession{
		Token:          uuid.NewString(),
		IdentityId:     identity.Id,
		Identity:       identity,
		LastActivityAt: time.Now(),
	}
	_, err := ctx.managers.ApiSession.Create(nil, entity, nil)
	ctx.NoError(err)
	return entity
}

func (ctx *TestContext) requireNewSession(apiSession *ApiSession, serviceId string, sessionType string) *Session {
	entity := &Session{
		Token:        uuid.NewString(),
		IdentityId:   apiSession.IdentityId,
		ApiSessionId: apiSession.Id,
		ServiceId:    serviceId,
		Type:         sessionType,
	}
	_, err := ctx.managers.Session.Create(entity, change.New())
	ctx.NoError(err)
	return entity
}

func (ctx *TestContext) requireNewServicePolicy(policyType string, identityRoles, serviceRoles []string) *ServicePolicy {
	policy := &ServicePolicy{
		Name:          eid.New(),
		Semantic:      persistence.SemanticAllOf,
		IdentityRoles: identityRoles,
		ServiceRoles:  serviceRoles,
		PolicyType:    policyType,
	}
	ctx.NoError(ctx.managers.ServicePolicy.Create(policy, change.New()))
	return policy
}

func (ctx *TestContext) requireNewEdgeRouterPolicy(identityRoles, edgeRouterRoles []string) *EdgeRouterPolicy {
	policy := &EdgeRouterPolicy{
		Name:            eid.New(),
		Semantic:        persistence.SemanticAllOf,
		IdentityRoles:   identityRoles,
		EdgeRouterRoles: edgeRouterRoles,
	}
	ctx.NoError(ctx.managers.EdgeRouterPolicy.Create(policy, change.New()))
	return policy
}

func (ctx *TestContext) requireNewServiceNewEdgeRouterPolicy(serviceRoles, edgeRouterRoles []string) *ServiceEdgeRouterPolicy {
	policy := &ServiceEdgeRouterPolicy{
		Name:            eid.New(),
		Semantic:        persistence.SemanticAllOf,
		ServiceRoles:    serviceRoles,
		EdgeRouterRoles: edgeRouterRoles,
	}
	ctx.NoError(ctx.managers.ServiceEdgeRouterPolicy.Create(policy, change.New()))
	return policy
}

func ss(vals ...string) []string {
	return vals
}
