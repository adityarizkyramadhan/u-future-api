package main

import (
	"fmt"
	"log"
	"os"
	"u-future-api/database/mysql"
	"u-future-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load env file\n cause %v", err)
	}

	db := mysql.InitDatabase()

	if db == nil {
		log.Fatal("init connection db failed")
	}
	err = mysql.Migrate()
	if err != nil {
		log.Fatal(err.Error())
	}

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.Use(middleware.CORS())

	router.Use(middleware.Timeout(30))

	router.GET("health", func(c *gin.Context) {
		c.String(200, "OK")
	})

	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	log.Printf("API run : port :%s\n", fmt.Sprintf(":%s", os.Getenv("PORT")))
}
