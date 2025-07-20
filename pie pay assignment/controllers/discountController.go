package controllers

import (
	"awesomeProject2/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DiscountController struct {
	Service interfaces.DiscountService
}

func NewDiscountController(service interfaces.DiscountService) *DiscountController {
	return &DiscountController{Service: service}
}

func (c *DiscountController) GetHighestDiscount(ctx *gin.Context) {
	amountToPayStr := ctx.Query("amountToPay")
	bankName := ctx.Query("bankName")
	paymentInstrument := ctx.Query("paymentInstrument") // optional

	_, err := strconv.ParseFloat(amountToPayStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amountToPay"})
		return
	}

	discount, err := c.Service.GetHighestDiscount(ctx, bankName, paymentInstrument)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"highestDiscountAmount": discount})
}
