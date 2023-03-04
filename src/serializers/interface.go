package serializers

import "context"

type Serializer[UrlParams any, Payload any, Model any, Response any] interface {
	ToModel(context.Context, Payload) (Model, error)
	ToResponse(context.Context, Model) (Response, error)
	BeforeSave(context.Context, Model) (Model, error)
	Save(context.Context, UrlParams, Model) (Model, error)
	AfterSave(context.Context, Model) error
}
