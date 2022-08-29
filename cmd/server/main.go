package main

import (
	"fmt"

	docs "github.com/brewerywiwis/food-review-backend/docs"
	"github.com/brewerywiwis/food-review-backend/pkg/domain/food"
	"github.com/brewerywiwis/food-review-backend/pkg/infrastructure/memory"
	"github.com/brewerywiwis/food-review-backend/pkg/interface/rest"
)

// @title Food Review System API
// @version 0.1.0
// @description This is a food review system api.
// @termsOfService http://swagger.io/terms/

// @contact.name brewerywiwis
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description the token is used to authenticate and authorize an user
func main() {
	foodRepository := memory.FoodStorage{}
	foodService := food.NewService(&foodRepository)
	server := rest.NewServer(foodService)

	// prepared to make a config env
	port := 8090
	basePath := "/"
	appVersion := "0.1.0"

	docs.SwaggerInfo.Host = fmt.Sprintf("http://localhost:%d", port)
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Version = appVersion

	server.Run(fmt.Sprintf(":%d", port))
}
