package main

import (
	"go-web/handler"
	"go-web/helper"
	"go-web/user"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	dsn := "root:qwertyuiop@tcp(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group(os.Getenv("API_VERSION"))

	api.POST("/users", userHandler.RegisterUser)
	api.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, helper.RandomString(20))
	})

	router.Run()
}
