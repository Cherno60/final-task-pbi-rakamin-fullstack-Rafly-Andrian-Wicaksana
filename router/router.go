package router

import (
	"github.com/gin-gonic/gin"
	controller "goAPI/controller"
	"net/http"
	"os"
)

func Routes() {
	r := gin.Default()
	r.GET("/home", Hello)
	//Users
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/logout", controller.Login)
	r.PUT("/logout", controller.Login)

	//Photos
	r.GET("/photos/", controller.PhotosIndex)
	r.POST("/photos/add", controller.PhotosIndex)
	r.PUT("/photos/edit/{id}", controller.PhotosIndex)
	r.DELETE("/photos", controller.PhotosIndex)
	r.Run("localhost:" + os.Getenv("PORT"))
}

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Welcome to home page")
}

//func Login(c *gin.Context) {
//	var payload loginPayload
//	if err := c.ShouldBindJSON(&payload); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	admin := "Unauthorized"
//	if payload.Username == "admin" && payload.Password == "123456" {
//		admin = "Welcome Admin"
//	}
//	c.JSON(http.StatusOK, gin.H{"message": admin})
//}
