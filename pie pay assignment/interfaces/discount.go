package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
)

type DiscountController interface {
	GetHighestDiscount(ctx *gin.Context)
}
type DiscountService interface {
	GetHighestDiscount(ctx context.Context, bankName, paymentInstrument string) (float64, error)
}
