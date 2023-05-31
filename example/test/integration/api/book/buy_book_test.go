package book

import (
	"net/http"
	"testing"

	"github.com/yagobatista/taco-go-web-framework/example/cmd/setup"
	"github.com/yagobatista/taco-go-web-framework/src/suite"
)

func TestBuyBookSuite(t *testing.T) {
	suite.Run(t, &BuyBookSuite{})
}

type BuyBookSuite struct {
	suite.ModelIntegrationSuite
}

func (this *BuyBookSuite) SetupTest() {
	this.SetServerConfig(setup.GetServerConfig())
	this.ModelIntegrationSuite.SetupTest()
}

func (this *BuyBookSuite) TestDisabledFeature() {
	statusCode := this.Client.Post("/internal/book/2/buy")

	this.Equal(http.StatusInternalServerError, statusCode)
}
