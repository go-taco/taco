package handlers

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm/clause"
)

type filterExempla struct {
	ColumnNameIN []string
	MyName       string
}

func TestGetFiltersSuite(t *testing.T) {
	suite.Run(t, &GetFiltersSuite{})
}

type GetFiltersSuite struct {
	suite.Suite
}

func (this *GetFiltersSuite) TestComplete() {
	expectedNames := []string{"one", "two"}

	filter := filterExempla{
		ColumnNameIN: expectedNames,
		MyName:       "name",
	}

	expression := getFilters(filter)

	expressions := expression.(clause.AndConditions).Exprs

	this.Require().Len(expressions, 2)

	inExpression := expressions[0].(clause.IN)

	this.Equal("column_name", inExpression.Column)
	this.Equal([]any{"one", "two"}, inExpression.Values)

	eqExpression := expressions[1].(clause.Eq)

	this.Equal("my_name", eqExpression.Column)
	this.Equal("name", eqExpression.Value)
}

func (this *GetFiltersSuite) TestGetINFiltersWrongSuffix() {
	expectedNames := []string{"one", "two"}

	_, err := getINFilters("nameIn", expectedNames)

	this.Error(err)
}
