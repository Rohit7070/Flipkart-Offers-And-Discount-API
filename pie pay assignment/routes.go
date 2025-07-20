package main

import (
	"awesomeProject2/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, offerController *controllers.OfferController, discountController *controllers.DiscountController) {
	r.POST("/offer", offerController.PostOffer)
	r.GET("/highest-discount", discountController.GetHighestDiscount)
}
