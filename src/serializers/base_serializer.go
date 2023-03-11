package serializers

import (
	"context"
)

type baseModelSerializer[Payload any, Model any, Response any] struct{}

func (this baseModelSerializer[Payload, Model, Response]) ToModel(context.Context, Payload) (model Model, err error) {
	panic("to be implemented on child")
}

func (this baseModelSerializer[Payload, Model, Response]) ToResponse(context.Context, Model) (resp Response, err error) {
	panic("to be implemented on child")
}
