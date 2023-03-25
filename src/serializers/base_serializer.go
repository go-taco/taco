package serializers

import (
	"context"
)

type baseModelSerializer[CreatePayload any, UpdatePayload any, Model any, Response any] struct{}

func (this baseModelSerializer[CreatePayload, UpdatePayload, Model, Response]) CreateToModel(context.Context, CreatePayload) (model Model, err error) {
	panic("to be implemented on child")
}

func (this baseModelSerializer[CreatePayload, UpdatePayload, Model, Response]) UpdateToModel(context.Context, UpdatePayload) (model Model, err error) {
	panic("to be implemented on child")
}

func (this baseModelSerializer[CreatePayload, UpdatePayload, Model, Response]) ToResponse(context.Context, Model) (resp Response, err error) {
	panic("to be implemented on child")
}
