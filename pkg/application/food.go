package application

import (
	"github.com/brewerywiwis/food-review-backend/pkg/domain/food"
	"github.com/gin-gonic/gin"
)

// FoodController provides to get the web request from router
type FoodController struct {
	service food.Service
}

// GetAll returns a list of Food
func (c *FoodController) GetAll(ctx *gin.Context) []food.Food {
	return c.service.ListAllFood()
}

// func (c *FoodController) CreateFood(ctx *gin.Context) (food.Food, error) {

// }
