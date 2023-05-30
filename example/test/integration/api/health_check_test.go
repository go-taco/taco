package health_check

import (
	"net/http"
	"testing"

	"github.com/yagobatista/taco-go-web-framework/example/cmd/setup"
	"github.com/yagobatista/taco-go-web-framework/src/suite"
)

func TestHealthCheckSuite(t *testing.T) {
	suite.Run(t, &HealthCheckSuite{})
}

type HealthCheckSuite struct {
	suite.IntegrationSuite
}

func (this *HealthCheckSuite) SetupTest() {
	this.SetServerConfig(setup.GetServerConfig())
	this.IntegrationSuite.SetupTest()
}

func (this *HealthCheckSuite) TestPing() {
	statusCode := this.Client.Get("/api/health/ping")

	this.Equal(http.StatusOK, statusCode)
}
