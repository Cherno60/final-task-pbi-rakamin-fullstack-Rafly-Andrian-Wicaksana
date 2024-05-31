package controller

import (
	"github.com/gin-gonic/gin"
	helper "goAPI/helper"
	models "goAPI/models"
	"net/http"
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
	// If there is no error
	// Save to DB or perform other actions here
	//fmt.Println("All OK")

	c.JSON(http.StatusOK, Response{
		Status:   http.StatusOK,
		Messages: "Validation Success",
		Errors:   nil,
	})
}
