package serializers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/handlers"
	"github.com/yagobatista/taco-go-web-framework/src/model"
)

type ToProcessor[Filter any, Payload any, Model any, Response any] func(serializer CreateSerializerInterface[Payload, Model, Response]) handlers.Processor[model.ModelUrlParams, Payload, Response]

func SerializerToCreateProcessor[Payload any, Model any, Response any](serializer CreateSerializerInterface[Payload, Model, Response]) handlers.Processor[model.ModelUrlParams, Payload, Response] {
	return func(ctx context.Context, urlParams model.ModelUrlParams, payload Payload) (resp Response, err error) {
		instance, err := serializer.CreateToModel(ctx, payload)
		if err != nil {
			return
		}

		instance, err = serializer.BeforeCreate(ctx, instance)
		if err != nil {
			return
		}

		instance, err = serializer.Create(ctx, urlParams, instance)
		if err != nil {
			return
		}

		err = serializer.AfterCreate(ctx, instance)
		if err != nil {
			return
		}

		return serializer.ToResponse(ctx, instance)
	}
}

func SerializerToUpdateProcessor[Payload any, Model any, Response any](serializer UpdateSerializerInterface[Payload, Model, Response]) handlers.Processor[model.ModelUrlParams, Payload, Response] {
	return func(ctx context.Context, urlParams model.ModelUrlParams, payload Payload) (resp Response, err error) {
		instance, err := serializer.UpdateToModel(ctx, payload)
		if err != nil {
			return
		}

		instance, err = serializer.BeforeUpdate(ctx, instance)
		if err != nil {
			return
		}

		instance, err = serializer.Update(ctx, urlParams, instance)
		if err != nil {
			return
		}

		err = serializer.AfterUpdate(ctx, instance)
		if err != nil {
			return
		}

		return serializer.ToResponse(ctx, instance)
	}
}

func SerializerToDetailProcessor[Filter any, Model any, Response any](serializer DetailSerializerInterface[Filter, Model, Response]) handlers.Processor[model.ModelUrlParams, Filter, Response] {
	return func(ctx context.Context, urlParams model.ModelUrlParams, filter Filter) (response Response, err error) {

		data, err := serializer.Detail(ctx, urlParams, filter)
		if err != nil {
			return response, err
		}

		return serializer.ToResponse(ctx, data)
	}
}

func SerializerToListProcessor[Filter any, Model any, Response any](serializer ListSerializerInterface[Filter, Model, Response]) handlers.Processor[model.ModelUrlParams, Filter, []Response] {
	return func(ctx context.Context, urlParams model.ModelUrlParams, filter Filter) (response []Response, err error) {
		data, err := serializer.List(ctx, filter)
		if err != nil {
			return response, err
		}

		response = make([]Response, len(data))

		for index, item := range data {
			resp, err := serializer.ToResponse(ctx, item)
			if err != nil {
				return response, err
			}

			response[index] = resp
		}

		return
	}
}
