package suite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"github.com/yagobatista/taco-go-web-framework/src/server"
)

type Client struct {
	*suite.Suite
	app *fiber.App

	username string
	password string
}

func NewClient(server *server.Server, suite *suite.Suite) Client {
	return Client{
		app:   server.GetFiberApp(),
		Suite: suite,
	}
}

func (this *Client) SetBasicAuth(username, password string) {
	this.username = username
	this.password = password
}

func (this *Client) Get(endpoint string) (statusCode int) {
	req := httptest.NewRequest(http.MethodGet, endpoint, nil)

	resp, err := this.app.Test(req, 3213213213)
	this.Require().NoError(err, "request failed")

	return resp.StatusCode
}

func (this *Client) Post(body any, endpoint string) (statusCode int, response any) {
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(body)
	this.Require().NoError(err, "failed to encode body")

	req := httptest.NewRequest(http.MethodPost, endpoint, &buf)

	if this.username != "" || this.password != "" {
		req.SetBasicAuth(this.username, this.password)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := this.app.Test(req, 3213213213)
	this.Require().NoError(err, "request failed")

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&response)

	return resp.StatusCode, response
}

func (this *Client) Postf(body any, endpoint string, args ...any) (statusCode int, response any) {
	return this.Post(body, fmt.Sprintf(endpoint, args...))
}
