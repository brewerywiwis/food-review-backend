package main

import (
	"github.com/brewerywiwis/food-review-backend/pkg/domain/food"
	"github.com/brewerywiwis/food-review-backend/pkg/infrastructure/memory"
	"github.com/brewerywiwis/food-review-backend/pkg/interface/rest"
)

func main() {
	foodRepository := memory.FoodStorage{}
	foodService := food.NewService(&foodRepository)
	server := rest.NewServer(foodService)

	server.Run(":8090")
}
