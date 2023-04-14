package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/invopop/jsonschema"
	"github.com/yagobatista/taco-go-web-framework/src/docs"
)

func (this *Handler[UrlParams, Payload, Response]) Docs(c *fiber.Ctx) error {
	var payload Payload
	var response Response

	expectedPayload, err := jsonschema.Reflect(&payload).MarshalJSON()
	if err != nil {
		return err
	}

	expectedResponse, err := jsonschema.Reflect(&response).MarshalJSON()
	if err != nil {
		return err
	}

	docsUrl := fmt.Sprintf("%s/docs", this.DocUrl)

	return docs.RenderDetailDoc(
		c,
		fiber.Map{
			"Payload":  string(expectedPayload),
			"Response": string(expectedResponse),
			"Title":    fmt.Sprintf("%s - %s", this.Method, this.Name),
			"Route":    strings.Replace(c.Route().Path, docsUrl, "", 1),
		},
	)
}
