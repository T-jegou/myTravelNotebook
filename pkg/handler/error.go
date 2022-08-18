package handler

import (
	"fmt"

	"github.com/T-jegou/myTravelNotebook/swagger_gen/models"
)

// Error is the handler error
type Error struct {
	StatusCode int
	Message    string
	Values     []interface{}
}

// ErrorMessage generates error messages
func ErrorMessage(s string, data ...interface{}) *models.Error {
	message := fmt.Sprintf(s, data...)
	return &models.Error{
		Message: &message,
	}
}
