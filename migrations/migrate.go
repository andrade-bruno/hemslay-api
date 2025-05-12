package main

import (
	"fmt"

	"github.com/andrade-bruno/hemslay-api/initializers"
	"github.com/andrade-bruno/hemslay-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	fmt.Println("Migrating database...")

	initializers.DB.AutoMigrate(&models.Post{})

	fmt.Println("Database migrated successfully!")
}
