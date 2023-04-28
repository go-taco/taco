package book

import "context"

type BuyBookUrlParams struct {
	ID uint `params:""id`
}

type BuyBookPayload struct {
	Quantity uint `json:"quantity"`
}

func (this BookHandler) BuyBook(ctx context.Context, urlParams BuyBookUrlParams, payload BuyBookPayload) (struct{}, error) {
	return struct{}{}, nil
}
