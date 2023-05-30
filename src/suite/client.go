package suite

import (
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"github.com/yagobatista/taco-go-web-framework/src/server"
)

type Client struct {
	suite.Suite
	app *fiber.App
}

func NewClient(server *server.Server, suite *suite.Suite) Client {
	return Client{
		app:   server.GetFiberApp(),
		Suite: *suite,
	}
}

func (this *Client) Get(endpoint string) (statusCode int) {

	req := httptest.NewRequest("GET", endpoint, nil)

	resp, err := this.app.Test(req)
	this.Require().NoError(err, "request failed")

	return resp.StatusCode
}
