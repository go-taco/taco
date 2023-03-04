package serializers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/model"
)

type ModelSerializer[UrlParams any, Payload any, Model any, Response any] struct{}

func (this ModelSerializer[UrlParams, Payload, Model, Response]) ToModel(context.Context, Payload) (model Model, err error) {
	panic("to be implemented on child")
}

func (this ModelSerializer[UrlParams, Payload, Model, Response]) ToResponse(context.Context, Model) (resp Response, err error) {
	panic("to be implemented on child")
}

func (this ModelSerializer[UrlParams, Payload, Model, Response]) BeforeSave(ctx context.Context, model Model) (Model, error) {
	return model, nil
}

func (this ModelSerializer[UrlParams, Payload, Model, Response]) Save(ctx context.Context, url UrlParams, instance Model) (Model, error) {
	return model.CreateModel(ctx, struct{}{}, instance)
}

func (this ModelSerializer[UrlParams, Payload, Model, Response]) AfterSave(context.Context, Model) error {
	return nil
}
