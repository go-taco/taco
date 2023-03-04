package serializers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/model"
)

type CreateModelSerializer[Payload any, Model any, Response any] struct {
	modelSerializer[Payload, Model, Response]
}

func (this CreateModelSerializer[Payload, Model, Response]) Save(ctx context.Context, urlParams model.ModelUrlParams, instance Model) (Model, error) {
	return model.CreateModel(ctx, struct{}{}, instance)
}

type UpdateModelSerializer[Payload any, Model any, Response any] struct {
	modelSerializer[Payload, Model, Response]
}

func (this UpdateModelSerializer[Payload, Model, Response]) Save(ctx context.Context, urlParams model.ModelUrlParams, instance Model) (Model, error) {
	return model.UpdateModel(ctx, urlParams, instance)
}

type GetModelSerializer[Payload any, Model any, Response any] struct {
	modelSerializer[Payload, Model, Response]
}

func (this GetModelSerializer[Payload, Model, Response]) Save(ctx context.Context, urlParams model.ModelUrlParams, instance Model) (Model, error) {
	return model.GetModel[Model](ctx, urlParams, struct{}{})
}
