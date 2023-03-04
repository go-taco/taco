package serializers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/model"
)

type Serializer[Payload any, Model any, Response any] interface {
	ToModel(context.Context, Payload) (Model, error)
	ToResponse(context.Context, Model) (Response, error)
	BeforeSave(context.Context, Model) (Model, error)
	Save(context.Context, model.ModelUrlParams, Model) (Model, error)
	AfterSave(context.Context, Model) error
}
