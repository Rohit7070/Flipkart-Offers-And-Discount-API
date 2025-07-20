package main

import (
	"awesomeProject2/config"
	"awesomeProject2/controllers"
	"awesomeProject2/repositories"
	"awesomeProject2/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	
	offerRepo := repositories.NewOfferRepository()
	offerService := services.NewOfferService(offerRepo)
	discountService := services.NewDiscountService(offerRepo)

	offerController := controllers.NewOfferController(offerService)
	discountController := controllers.NewDiscountController(discountService)

	r := gin.Default()
	RegisterRoutes(r, offerController, discountController)
	r.Run(":8080")
}
