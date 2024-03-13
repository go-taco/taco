package suite

import "github.com/yagobatista/taco-go-web-framework/src/database"

type ModelIntegrationSuite struct {
	IntegrationSuite
}

func (this *ModelIntegrationSuite) Create(model any) {
	err := database.GetConnectionFromCtx(this.Ctx).Create(model).Error
	this.Require().NoError(err)
}
