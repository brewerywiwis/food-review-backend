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
