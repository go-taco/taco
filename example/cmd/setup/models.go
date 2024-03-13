package setup

import "github.com/yagobatista/taco-go-web-framework/example/models"

func GetModelsRegistry() []any {
	return []any{
		models.Book{},
	}
}
