package setup

import "github.com/yagobatista/taco-go-web-framework/example/structs"

func GetTables() []any {
	return []any{
		structs.Book{},
	}
}
