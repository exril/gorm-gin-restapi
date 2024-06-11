package controllers

import (
	"4mount/backend/initializers"
	"4mount/backend/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: "Hello", Body: "Post Body"}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}


	c.JSON(200, gin.H{
		"post" : post,
	})
}