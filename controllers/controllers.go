package controllers

import (
	"github.com/gaurav/go-crud/initializers"
	"github.com/gaurav/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//get data of body
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})

	//create post

	//return it
}

func PostIndex(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostShow(c *gin.Context) {
	//get id off url
	var post models.Post

	id := c.Param("id")

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostsUpdate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	result := initializers.DB.Model(&post).Updates(
		models.Post{
			Title: body.Title,
			Body:  body.Body,
		})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//get id off url
	var post models.Post

	id := c.Param("id")

	initializers.DB.Delete(&post, id)

	c.Status(200)

}
