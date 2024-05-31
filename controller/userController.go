package controller

import (
	"github.com/gin-gonic/gin"
	"goAPI/database"
	helper "goAPI/helper"
	models "goAPI/models"
	"net/http"
	"time"
)

type Response struct {
	Status   int         `json:"status"`
	Messages string      `json:"messages"`
	Errors   interface{} `json:"errors"`
}

func Login(c *gin.Context) {

}
func Register(c *gin.Context) {
	user := new(models.Users)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:   http.StatusBadRequest,
			Messages: "Invalid request",
			Errors:   err.Error(),
		})
		return
	}

	errorList := helper.ValidateUser(user)
	if errorList != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:   http.StatusBadRequest,
			Messages: "Validation Error",
			Errors:   errorList,
		})
		return
	}

	userRegister := models.Users{
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := database.DB.Create(&userRegister)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:   http.StatusOK,
		Messages: "User Created",
		Errors:   nil,
	})
}
