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

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Status   int    `json:"status"`
	Messages string `json:"messages"`
}

func GetUserFromID(c *gin.Context) {

	uuid := c.Param("uuid")

	var user models.Users
	if err := database.DB.First(&user, "uuid = ?", uuid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Map the user data to the UserResponse struct
	userResponse := UserResponse{
		Email:    user.Email,
		Username: user.Username,
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":   http.StatusOK,
		"Messages": userResponse,
		"Errors":   nil,
	})
}

func Login(c *gin.Context) {
	var LoginRequest LoginRequest
	if err := c.ShouldBindJSON(&LoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var users []models.Users
	result := database.DB.Where("email = ?", LoginRequest.Email).Find(&users)
	if result.Error != nil {
		c.Status(400)
		return
	}
	if result.RowsAffected < 1 {
		c.JSON(http.StatusUnauthorized, LoginResponse{
			Status:   http.StatusUnauthorized,
			Messages: "Email or password is incorrect",
		})
		return
	}
	// Do auth
	c.JSON(http.StatusOK, LoginResponse{
		Status:   http.StatusOK,
		Messages: "Logined successfully!, Hello " + LoginRequest.Email + "!",
	})
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

func UserUpdate(c *gin.Context) {
	var UserEdit struct {
		Username string
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&UserEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get Parameter
	uuid := c.Param("uuid")

	// Find the data
	var user models.Users
	database.DB.First(&user, "uuid = ?", uuid)
	// Update
	database.DB.Model(&user).Updates(models.Users{
		Username: UserEdit.Username,
		Email:    UserEdit.Email,
		Password: UserEdit.Password,
	})

	c.JSON(http.StatusOK, gin.H{
		"Status":   http.StatusOK,
		"Messages": UserEdit,
		"Errors":   nil,
	})
}

func UserDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	var user models.Users
	if err := database.DB.First(&user, "uuid = ?", uuid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user, "uuid = ?", uuid)
	c.JSON(http.StatusOK, Response{
		Status:   http.StatusOK,
		Messages: "User Deleted",
		Errors:   nil,
	})
}
