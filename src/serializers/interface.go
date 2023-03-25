package serializers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/model"
)

type SerializerInterface[Model any, Response any] interface {
	ToResponse(context.Context, Model) (Response, error)
}

type CreateSerializerInterface[Payload any, Model any, Response any] interface {
	SerializerInterface[Model, Response]

	CreateToModel(context.Context, Payload) (Model, error)

	BeforeCreate(context.Context, Model) (Model, error)
	Create(context.Context, model.ModelUrlParams, Model) (Model, error)
	AfterCreate(context.Context, Model) error
}

type UpdateSerializerInterface[Payload any, Model any, Response any] interface {
	SerializerInterface[Model, Response]

	UpdateToModel(context.Context, Payload) (Model, error)

	BeforeUpdate(context.Context, Model) (Model, error)
	Update(context.Context, model.ModelUrlParams, Model) (Model, error)
	AfterUpdate(context.Context, Model) error
}

type DetailSerializerInterface[Filter any, Model any, Response any] interface {
	SerializerInterface[Model, Response]

	Detail(context.Context, model.ModelUrlParams, Filter) (Model, error)
}

type ListSerializerInterface[Filter any, Model any, Response any] interface {
	SerializerInterface[Model, Response]

	List(context.Context, Filter) ([]Model, error)
}

type ModelSerializerInterface[CreatePayload any, UpdatePayload any, DetailQueryParams any, Filter any, Model any, Response any] interface {
	CreateSerializerInterface[CreatePayload, Model, Response]
	UpdateSerializerInterface[UpdatePayload, Model, Response]
	DetailSerializerInterface[DetailQueryParams, Model, Response]
	ListSerializerInterface[Filter, Model, Response]
}
