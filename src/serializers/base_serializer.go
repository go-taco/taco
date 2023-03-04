package serializers

import "context"

type modelSerializer[Payload any, Model any, Response any] struct{}

func (this modelSerializer[Payload, Model, Response]) ToModel(context.Context, Payload) (model Model, err error) {
	panic("to be implemented on child")
}

func (this modelSerializer[Payload, Model, Response]) ToResponse(context.Context, Model) (resp Response, err error) {
	panic("to be implemented on child")
}

func (this modelSerializer[Payload, Model, Response]) BeforeSave(ctx context.Context, model Model) (Model, error) {
	return model, nil
}

func (this modelSerializer[Payload, Model, Response]) AfterSave(context.Context, Model) error {
	return nil
}
