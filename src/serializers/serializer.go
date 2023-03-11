package serializers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/model"
)

type CreateModelSerializer[Payload any, Model any, Response any] struct {
	baseModelSerializer[Payload, Model, Response]
}

func (this CreateModelSerializer[Payload, Model, Response]) Create(ctx context.Context, urlParams model.ModelUrlParams, instance Model) (Model, error) {
	return model.CreateModel(ctx, struct{}{}, instance)
}

func (this CreateModelSerializer[Payload, Model, Response]) BeforeCreate(ctx context.Context, model Model) (Model, error) {
	return model, nil
}

func (this CreateModelSerializer[Payload, Model, Response]) AfterCreate(context.Context, Model) error {
	return nil
}

type UpdateModelSerializer[Payload any, Model any, Response any] struct {
	baseModelSerializer[Payload, Model, Response]
}

func (this UpdateModelSerializer[Payload, Model, Response]) Update(ctx context.Context, urlParams model.ModelUrlParams, instance Model) (Model, error) {
	return model.UpdateModel(ctx, urlParams, instance)
}

func (this UpdateModelSerializer[Payload, Model, Response]) BeforeUpdate(ctx context.Context, model Model) (Model, error) {
	return model, nil
}

func (this UpdateModelSerializer[Payload, Model, Response]) AfterUpdate(context.Context, Model) error {
	return nil
}

type DetailModelSerializer[Payload any, Model any, Response any] struct {
	baseModelSerializer[Payload, Model, Response]
}

func (this DetailModelSerializer[Payload, Model, Response]) Detail(ctx context.Context, urlParams model.ModelUrlParams, instance Payload) (Model, error) {
	return model.GetModel[Model](ctx, urlParams, struct{}{})
}

type ListModelSerializer[Payload any, Model any, Response any] struct {
	baseModelSerializer[Payload, Model, Response]
}

func (this ListModelSerializer[Payload, Model, Response]) List(ctx context.Context, payload Payload) ([]Model, error) {
	return model.ListModel[Payload, Model](ctx, struct{}{}, payload)
}

type ModelSerializer[Payload any, Model any, Response any] struct {
	baseModelSerializer[Payload, Model, Response]
	CreateModelSerializer[Payload, Model, Response]
	UpdateModelSerializer[Payload, Model, Response]
	DetailModelSerializer[Payload, Model, Response]
	ListModelSerializer[Payload, Model, Response]
}
