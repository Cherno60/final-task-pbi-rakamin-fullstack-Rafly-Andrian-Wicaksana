package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	models "goAPI/Models"
	"reflect"
)

func UserValidation(c *gin.Context) error {
	user := new(models.Users)
	if err := c.ShouldBind(&user); err != nil {
		return
	}
	validate := *validator.New()
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
				errMsg = fmt.Sprintf("%s is invalid email", fieldName)
			case "min":
				errMsg = fmt.Sprintf("%s is too short", fieldName)
			}
			errorList[fieldName] = errMsg
		}

	}
	return nil
}
