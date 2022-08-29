package rest

import (
	"github.com/brewerywiwis/food-review-backend/pkg/domain/food"
)

type ListAllFoodResponse struct {
	Data []food.Food `json:"data"`
}
type CreateFoodResponse struct {
	Message string `json:"message"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse is use to create a new error response
func NewErrorResponse(code int) ErrorResponse {
	var msg string
	switch code {
	case 0:
		msg = "Body is invalidated"
	case 1:
		msg = "Something is error"
	default:
		msg = "Body is invalidated"
	}
	return ErrorResponse{msg}
}
