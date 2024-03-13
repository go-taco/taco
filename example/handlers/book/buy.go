package book

import (
	"context"
	"errors"

	project_configs "github.com/yagobatista/taco-go-web-framework/example/configs"
	"github.com/yagobatista/taco-go-web-framework/example/models"
	"github.com/yagobatista/taco-go-web-framework/src/configs"
	"github.com/yagobatista/taco-go-web-framework/src/database"
)

type BuyBookUrlParams struct {
	ID uint `params:"id"`
}

type BuyBookPayload struct {
	Quantity uint64 `json:"quantity"`
}

func (this BookHandler) BuyBook(ctx context.Context, urlParams BuyBookUrlParams, payload BuyBookPayload) (struct{}, error) {
	cfg := configs.GetFromCtx[project_configs.Configs](ctx)

	if cfg.DisableBuyFeature {
		return struct{}{}, errors.New("feature disabled")
	}

	var instance models.Book
	var filter models.Book
	filter.ID = urlParams.ID

	err := database.GetConnectionFromCtx(ctx).Where(filter).Find(&instance).Error
	if err != nil {
		return struct{}{}, err
	}

	instance.AvailableCopies -= payload.Quantity

	if int64(instance.AvailableCopies) < 0 {
		return struct{}{}, errors.New("invalid number of copies")
	}

	return struct{}{}, nil
}
