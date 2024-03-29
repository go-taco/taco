package model

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/database"
)

type ModelUrlParams struct {
	ID string `params:"id"`
}

func CreateModel[Model any](ctx context.Context, urlParams struct{}, payload Model) (Model, error) {
	err := database.GetConnectionFromCtx(ctx).
		Create(&payload).
		Error

	return payload, err
}

func UpdateModel[Model any](ctx context.Context, urlParams ModelUrlParams, payload Model) (Model, error) {
	var model Model

	conn := database.GetConnectionFromCtx(ctx)

	err := conn.
		Model(model).
		Where("id = ? ", urlParams.ID).
		UpdateColumns(&payload).
		Error
	if err != nil {
		return payload, err
	}

	var instance Model

	err = conn.
		Model(model).
		Where("id = ? ", urlParams.ID).
		Find(&instance).
		Error
	if err != nil {
		return payload, err
	}

	return instance, nil
}

func GetModel[Model any](ctx context.Context, urlParams ModelUrlParams, payload struct{}) (Model, error) {
	var model Model
	var instance Model

	err := database.GetConnectionFromCtx(ctx).
		Model(model).
		Where("id = ?", urlParams.ID).
		Find(&instance).
		Error

	return instance, err
}
