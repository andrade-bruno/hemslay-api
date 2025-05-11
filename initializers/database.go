package initializers

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectDb() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := url.QueryEscape(os.Getenv("DB_PASSWORD"))

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s", user, password, host, port, database)

	_, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	fmt.Println("Database connected!")
}
