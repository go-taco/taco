package model

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/gobeam/stringy"
	"github.com/yagobatista/taco-go-web-framework/src/server"
	"gorm.io/gorm/clause"
)

func ListModel[Filters any, Model any](ctx context.Context, urlParams struct{}, filters Filters) (results []Model, err error) {
	var model Model

	err = server.GetConnectionFromCtx(ctx).
		Model(model).
		Where(getFilters(filters)).
		Find(&results).
		Error

	return results, err
}

func getFilters(filters any) clause.Expression {

	v := reflect.ValueOf(filters)

	var expressions []clause.Expression

	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name

		field := v.Field(i)
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		value := field.Interface()

		exp, err := getINFilters(name, value)
		if err == nil {
			expressions = append(expressions, *exp)
			continue
		}

		expressions = append(expressions, clause.Eq{
			Column: getColumnName(name),
			Value:  value,
		})
	}

	return clause.And(
		expressions...,
	)
}

func getINFilters(fieldName string, value any) (*clause.IN, error) {
	if !strings.HasSuffix(fieldName, "IN") {
		return nil, errors.New("not a IN type")
	}

	reflectValue := reflect.Indirect(reflect.ValueOf(value))

	valueLen := reflectValue.Len()
	values := make([]interface{}, valueLen)
	for i := 0; i < valueLen; i++ {
		values[i] = reflectValue.Index(i).Interface()
	}

	name := strings.Replace(fieldName, "IN", "", 1)

	return &clause.IN{
		Column: getColumnName(name),
		Values: values,
	}, nil
}

func getColumnName(column string) string {
	return stringy.New(column).SnakeCase().ToLower()
}
