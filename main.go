package main

import (
	"github.com/gaurav/go-crud/controllers"
	"github.com/gaurav/go-crud/initializers"
	"github.com/gaurav/go-crud/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
	initializers.MigrateToDb()

}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.POST("/createUser", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)
	r.Run()
}
