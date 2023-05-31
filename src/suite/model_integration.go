package suite

import (
	"github.com/yagobatista/taco-go-web-framework/src/database"
	"gorm.io/gorm"
)

type ModelIntegrationSuite struct {
	IntegrationSuite

	Conn *gorm.DB
}

func (this *ModelIntegrationSuite) SetupTest() {
	this.IntegrationSuite.SetupTest()
	this.Conn = database.GetConnectionFromCtx(this.Ctx)
}
