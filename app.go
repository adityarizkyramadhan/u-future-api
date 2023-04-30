package main

import (
	"fmt"
	"log"
	"os"
	rpSchool "u-future-api/api/school/repository"
	ucSchool "u-future-api/api/school/usecase"
	rpStudent "u-future-api/api/student/repository"
	ucStudent "u-future-api/api/student/usecase"
	"u-future-api/database/mysql"
	"u-future-api/middleware"
	"u-future-api/models"

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
	err = mysql.Migrate(new(models.School), new(models.Student))
	if err != nil {
		log.Fatal(err.Error())
	}

	repoSchool := rpSchool.New(db)
	useCaseSchool := ucSchool.New(repoSchool)

	repoStudent := rpStudent.New(db)
	useCaseStudent := ucStudent.New(repoStudent)

	// Generate Faker School
	err = useCaseSchool.GenerateFaker()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Generate Faker Student
	err = useCaseStudent.GenerateFaker(useCaseSchool)
	if err != nil {
		log.Fatalln(err.Error())
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
