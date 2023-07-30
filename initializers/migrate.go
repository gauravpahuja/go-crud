package initializers

import (
	"github.com/gaurav/go-crud/models"
)

func init() {
	LoadEnvVariable()
	ConnectToDb()
}

func MigrateToDb() {
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.User{})
}
