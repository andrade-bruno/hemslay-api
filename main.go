package main

import (
	"fmt"

	"github.com/andrade-bruno/hemslay-api/controllers"
	"github.com/andrade-bruno/hemslay-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	router := gin.Default()

	router.GET("/healthcheck", controllers.HealthCheck)

	router.GET("/posts", controllers.IndexPosts)
	router.GET("/posts/:id", controllers.GetPost)
	router.POST("/posts", controllers.CreatePost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	fmt.Println("Hemslay API is up and running!")
}
