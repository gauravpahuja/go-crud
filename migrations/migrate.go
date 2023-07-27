package main

import (
	"github.com/gaurav/go-crud/initializers"

	"github.com/gaurav/go-crud/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
