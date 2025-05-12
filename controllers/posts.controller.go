package controllers

import (
	"fmt"
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

	c.JSON(http.StatusCreated, gin.H{
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

func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	output := initializers.DB.First(&post, id)

	if output.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Post with ID %s not found", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	output := initializers.DB.First(&post, id)

	if output.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Post with ID %s not found", id),
		})
		return
	}

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	output := initializers.DB.First(&post, id)

	if output.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Post #%s not found", id),
		})
		return
	}

	initializers.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Post #%s deleted", id),
	})
}
