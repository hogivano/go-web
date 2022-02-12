package main

import (
	"fmt"
	"go-web/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:qwertyuiop@tcp(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(db)

	router := gin.Default()
	router.GET("/db", handler)

	router.Run()
}

func handler(c *gin.Context) {
	dsn := "root:qwertyuiop@tcp(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User

	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
