package helper

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	models "goAPI/Models"
)

type Response struct {
	Status   int         `json:"status"`
	Messages string      `json:"messages"`
	Errors   interface{} `json:"errors"`
}

// ValidateUser validates the user struct
func ValidateUser(user *models.Users) map[string]string {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		errorList := make(map[string]string)

		for _, e := range errors {
			var errMsg string
			field, _ := reflect.TypeOf(*user).FieldByName(e.StructField())
			fieldName := field.Tag.Get("json")

			switch e.Tag() {
			case "required":
				errMsg = fmt.Sprintf("%s is required", fieldName)
			case "email":
				errMsg = fmt.Sprintf("%s is an invalid email", fieldName)
			case "min":
				errMsg = fmt.Sprintf("%s is too short, minimum length is %s", fieldName, e.Param())
			default:
				errMsg = fmt.Sprintf("%s is invalid", fieldName)
			}
			errorList[fieldName] = errMsg
		}
		return errorList
	}
	return nil
}

// UserRegistration handles user registration
func UserRegistration(c *gin.Context) {
	user := new(models.Users)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:   http.StatusBadRequest,
			Messages: "Invalid request",
			Errors:   err.Error(),
		})
		return
	}

	errorList := ValidateUser(user)
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

	c.JSON(http.StatusOK, Response{
		Status:   http.StatusOK,
		Messages: "Validation Success",
		Errors:   nil,
	})
}
