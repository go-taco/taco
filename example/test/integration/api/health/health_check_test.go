package health_check

import (
	"net/http"
	"testing"

	"github.com/yagobatista/taco-go-web-framework/example/cmd/setup"
	"github.com/yagobatista/taco-go-web-framework/src/database"
	"github.com/yagobatista/taco-go-web-framework/src/suite"
)

func TestHealthCheckSuite(t *testing.T) {
	suite.Run(t, &HealthCheckSuite{})
}

type HealthCheckSuite struct {
	suite.ModelIntegrationSuite
}

func (this *HealthCheckSuite) SetupTest() {
	this.SetServerConfig(setup.GetServerConfig())
	this.ModelIntegrationSuite.SetupTest()
}

func (this *HealthCheckSuite) TestPing() {
	statusCode := this.Client.Get("/api/health/ping")

	this.Equal(http.StatusOK, statusCode)
}

func (this *HealthCheckSuite) TestClosedConnection() {
	db, err := database.GetConnectionFromCtx(this.Ctx).DB()
	this.Require().NoError(err, "failed to get DB")

	err = db.Close()
	this.Require().NoError(err, "failed to close connection")

	statusCode := this.Client.Get("/api/health/ping")

	this.Equal(http.StatusInternalServerError, statusCode)
}
