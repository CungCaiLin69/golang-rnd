package middleware

import (
	"fmt"
	"golang-rnd/lib"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateBody[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload T

		if err := c.ShouldBindJSON(&payload); err != nil {
			fmt.Println("❌ JSON Bind Error:", err) // Debugging log
			apiErr := lib.NewApiError("Invalid JSON", []string{err.Error()}, false, nil)
			c.JSON(http.StatusBadRequest, apiErr)
			c.Abort()
			return
		}

		if err := validate.Struct(payload); err != nil {
			errors := make(map[string]interface{})
			rawErrors := []string{}
			t := reflect.TypeOf(payload)

			for _, err := range err.(validator.ValidationErrors) {
				field, _ := t.FieldByName(err.StructField()) // Get field name from struct
				errors[field.Name] = fmt.Sprintf("is %s", err.Tag())
				rawErrors = append(rawErrors, fmt.Sprintf("%s is %s", err.Field(), err.ActualTag()))
			}

			apiErr := lib.NewApiError("Validation Error", rawErrors, false, errors)
			c.JSON(http.StatusBadRequest, apiErr)
			c.Abort()
			return
		}

		c.Set("payload", payload)
		c.Next()
	}
}
