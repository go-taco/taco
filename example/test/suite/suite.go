package suite

import (
	"testing"

	"github.com/yagobatista/taco-go-web-framework/example/cmd/setup"
	"github.com/yagobatista/taco-go-web-framework/src/suite"
)

func Run(t *testing.T, s suite.TestingSuite) {
	suite.Run(t, s)
}

type ModelIntegrationSuite struct {
	suite.ModelIntegrationSuite
}

func (this *ModelIntegrationSuite) SetupSuite() {
	this.SetTables(setup.GetModelsRegistry())
	this.ModelIntegrationSuite.SetupSuite()
}

func (this *ModelIntegrationSuite) SetupTest() {
	this.SetServerConfig(setup.GetServerConfig())
	this.ModelIntegrationSuite.SetupTest()
}
