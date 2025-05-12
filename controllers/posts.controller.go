package controllers

import (
	"net/http"

	"github.com/andrade-bruno/hemslay-api/initializers"
	"github.com/andrade-bruno/hemslay-api/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	output := initializers.DB.Create(&post)

	if output.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rowsAffected": output.RowsAffected,
		"data":         post,
	})
}

func IndexPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No posts found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}
