package serializers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/handlers"
)

func SerializerToProcessor[UrlParams any, Payload any, Model any, Response any](serializer Serializer[UrlParams, Payload, Model, Response]) handlers.Processor[UrlParams, Payload, Response] {
	return func(ctx context.Context, urlParams UrlParams, payload Payload) (resp Response, err error) {
		model, err := serializer.ToModel(ctx, payload)
		if err != nil {
			return
		}

		model, err = serializer.BeforeSave(ctx, model)
		if err != nil {
			return
		}

		model, err = serializer.Save(ctx, urlParams, model)
		if err != nil {
			return
		}

		err = serializer.AfterSave(ctx, model)
		if err != nil {
			return
		}

		return serializer.ToResponse(ctx, model)
	}
}
