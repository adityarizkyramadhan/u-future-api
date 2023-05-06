package main

import (
	"fmt"
	"log"
	"os"
	ctSchool "u-future-api/api/school/controller"
	rpSchool "u-future-api/api/school/repository"
	ucSchool "u-future-api/api/school/usecase"
	ctStudent "u-future-api/api/student/controller"
	rpStudent "u-future-api/api/student/repository"
	ucStudent "u-future-api/api/student/usecase"
	ucQuiz "u-future-api/api/quiz/usecase"
	rpQuiz "u-future-api/api/quiz/repository"
	ctQuiz "u-future-api/api/quiz/controller"
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
	err = mysql.Migrate(
		&models.School{},
		&models.Student{},
		&models.Quiz{},
		&models.Question{},
		&models.Option{},
		&models.QuizResult{},
	)
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

	repoSchool := rpSchool.New(db)
	useCaseSchool := ucSchool.New(repoSchool)

	repoStudent := rpStudent.New(db)
	useCaseStudent := ucStudent.New(repoStudent)

	repoQuiz := rpQuiz.New(db)
	useCaseQuiz := ucQuiz.New(repoQuiz)

	var schoolCount int64
	if err := db.Model(&models.School{}).Count(&schoolCount).Error; err != nil {
		log.Fatalln(err.Error())
	}
	if schoolCount == 0 {
		// Generate Faker School
		err = useCaseSchool.GenerateFaker()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	// check if the student table already has data
	var studentCount int64
	if err := db.Model(&models.Student{}).Count(&studentCount).Error; err != nil {
		log.Fatalln(err.Error())
	}
	if studentCount == 0 {
		// Generate Faker Student
		err = useCaseStudent.GenerateFaker(useCaseSchool)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	v1 := router.Group("api/v1")
	ctrlStudent := ctStudent.New(useCaseStudent)
	studentGrop := v1.Group("student")
	ctrlStudent.Mount(studentGrop)

	ctrlSchool := ctSchool.New(useCaseSchool)
	schoolGroup := v1.Group("school")
	ctrlSchool.Mount(schoolGroup)

	ctrlQuiz := ctQuiz.New(useCaseQuiz)
	quizGroup := v1.Group("quiz")
	ctrlQuiz.Mount(quizGroup)

	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	log.Printf("API run : port :%s\n", fmt.Sprintf(":%s", os.Getenv("PORT")))
}
