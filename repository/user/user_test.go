package user

import (
	"be/configs"
	"be/entities"
	"be/utils"
	"testing"

	"github.com/lithammer/shortuuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)
	var r = New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Product{})
	db.AutoMigrate(&entities.User{})

	t.Run("success run create", func(t *testing.T) {
		var mock1 = entities.User{UserName: shortuuid.New(), Email: shortuuid.New(), Password: shortuuid.New(), Name: shortuuid.New()}

		var res, err = r.Create(mock1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}
