package book

import (
	"net/http"
	"testing"

	"github.com/yagobatista/taco-go-web-framework/example/handlers/book"
	"github.com/yagobatista/taco-go-web-framework/example/structs"
	"github.com/yagobatista/taco-go-web-framework/example/test/suite"
)

func TestBuyBookSuite(t *testing.T) {
	suite.Run(t, &BuyBookSuite{})
}

type BuyBookSuite struct {
	suite.ModelIntegrationSuite
}

func (this *BuyBookSuite) SetupTest() {
	this.ModelIntegrationSuite.SetupTest()
	this.Client.SetBasicAuth("admin2", "admin2")
}

func (this *BuyBookSuite) TestValidCopies() {
	instance := structs.Book{
		Title:           "My book",
		Author:          "Me",
		AvailableCopies: 25,
	}
	this.Create(&instance)

	statusCode, _ := this.Client.Postf(
		book.BuyBookPayload{
			Quantity: 2,
		},
		"/internal/book/%d/buy", instance.ID,
	)

	this.Equal(http.StatusCreated, statusCode)
}

func (this *BuyBookSuite) TestBuyTooMuchCopies() {
	instance := structs.Book{
		Title:           "My book",
		Author:          "Me",
		AvailableCopies: 25,
	}
	this.Create(&instance)

	statusCode, _ := this.Client.Postf(
		book.BuyBookPayload{
			Quantity: 30,
		},
		"/internal/book/%d/buy", instance.ID,
	)

	this.Equal(http.StatusInternalServerError, statusCode)
}
