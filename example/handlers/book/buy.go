package book

import (
	"context"
	"errors"

	project_configs "github.com/yagobatista/taco-go-web-framework/example/configs"
	"github.com/yagobatista/taco-go-web-framework/src/configs"
)

type BuyBookUrlParams struct {
	ID uint `params:"id"`
}

type BuyBookPayload struct {
	Quantity uint `json:"quantity"`
}

func (this BookHandler) BuyBook(ctx context.Context, urlParams BuyBookUrlParams, payload BuyBookPayload) (struct{}, error) {
	cfg := configs.GetFromCtx[project_configs.Configs](ctx)

	if cfg.DisableBuyFeature {
		return struct{}{}, errors.New("feature disabled")
	}

	return struct{}{}, nil
}
