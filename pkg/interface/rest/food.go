package rest

import (
	"time"

	"github.com/brewerywiwis/food-review-backend/pkg/domain/food"
	"github.com/gin-gonic/gin"
)

type foodController struct {
	service food.Service
}

func (c *foodController) listAllFood(ctx *gin.Context) {
	response := ListAllFoodResponse{c.service.ListAllFood()}
	ctx.JSON(200, response)
}

func (c *foodController) createFood(ctx *gin.Context) {
	req := CreateFoodRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, NewErrorResponse(0))
	}
	food := food.Food{
		Name:        req.Name,
		Description: req.Description,
		Img:         req.Img,
		Location:    req.Location,
		Rate:        req.Rate,
		CreatedAt:   time.Now(),
	}
	if err := c.service.AddFood(food); err != nil {
		ctx.JSON(400, NewErrorResponse(0))
	}
	response := CreateFoodResponse{"success"}
	ctx.JSON(200, response)
}
