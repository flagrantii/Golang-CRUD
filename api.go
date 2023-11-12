package main

import (
	"go-crud/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/go_crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		var users []model.User
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		var user model.User
		db.First(&user, c.Param("id"))
		c.JSON(http.StatusOK, user)
	})
	r.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&user)
		c.JSON(http.StatusOK, user)
	})
	r.DELETE("/users/:id", func(c *gin.Context) {
		var user model.User
		db.Delete(&user, c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})
	r.PUT("/users/:id", func(c *gin.Context) {
		var user model.User
		db.First(&user, c.Param("id"))
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&user)
		c.JSON(http.StatusOK, user)
	})
	r.Use(cors.Default())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
