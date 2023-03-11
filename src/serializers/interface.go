package serializers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/model"
)

type SerializerInterface[Payload any, Model any, Response any] interface {
	ToModel(context.Context, Payload) (Model, error)
	serializeToResponseInterface[Model, Response]
}

type serializeToResponseInterface[Model any, Response any] interface {
	ToResponse(context.Context, Model) (Response, error)
}

type CreateSerializerInterface[Payload any, Model any, Response any] interface {
	SerializerInterface[Payload, Model, Response]

	BeforeCreate(context.Context, Model) (Model, error)
	Create(context.Context, model.ModelUrlParams, Model) (Model, error)
	AfterCreate(context.Context, Model) error
}

type UpdateSerializerInterface[Payload any, Model any, Response any] interface {
	SerializerInterface[Payload, Model, Response]

	BeforeUpdate(context.Context, Model) (Model, error)
	Update(context.Context, model.ModelUrlParams, Model) (Model, error)
	AfterUpdate(context.Context, Model) error
}

type DetailSerializerInterface[Filter any, Model any, Response any] interface {
	serializeToResponseInterface[Model, Response]

	Detail(context.Context, model.ModelUrlParams, Filter) (Model, error)
}

type ListSerializerInterface[Filter any, Model any, Response any] interface {
	serializeToResponseInterface[Model, Response]

	List(context.Context, Filter) ([]Model, error)
}

type ModelSerializerInterface[Filter any, Payload any, Model any, Response any] interface {
	CreateSerializerInterface[Payload, Model, Response]
	UpdateSerializerInterface[Payload, Model, Response]
	DetailSerializerInterface[Filter, Model, Response]
	ListSerializerInterface[Filter, Model, Response]
}
