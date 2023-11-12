package book

import (
	"net/http"
	"testing"

	"github.com/yagobatista/taco-go-web-framework/example/cmd/setup"
	"github.com/yagobatista/taco-go-web-framework/example/handlers/book"
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
	this.Client.SetBasicAuth("admin2", "admin2")
}

func (this *BuyBookSuite) TestDisabledFeature() {
	statusCode, _ := this.Client.Postf(
		book.BuyBookPayload{
			Quantity: 2,
		},
		"/internal/book/%d/buy", 2,
	)

	this.Equal(http.StatusInternalServerError, statusCode)
}
