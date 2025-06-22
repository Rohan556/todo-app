package loggers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	ErrorMessage string
}

func HandleResponse(ctx *gin.Context, status int, content any) {
	ctx.JSON(status, content)
}

func HandleValidationErrors(errors validator.ValidationErrors) ErrorResponse {
	message := ""
	for _, err := range errors {
		switch err.ActualTag() {
		case "required":
			message += fmt.Sprintf("%s is required, ", err.Field())
		}
	}

	return ErrorResponse{
		ErrorMessage: message,
	}
}

func ValidateRequestBody(ctx *gin.Context, requestBody any) bool {
	validate := validator.New()

	// Validate the User struct
	if err := validate.Struct(requestBody); err != nil {
		errors := err.(validator.ValidationErrors)
		errorMessage := HandleValidationErrors(errors)
		HandleResponse(ctx, http.StatusBadRequest, errorMessage)
		return false
	}

	return true
}
