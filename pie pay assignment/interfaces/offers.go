package interfaces

import (
	"awesomeProject2/models"
	"context"
	"github.com/gin-gonic/gin"
)

type OfferController interface {
	PostOffer(ctx *gin.Context)
}
type OffersService interface {
	SaveOffers(apiResp models.FlipkartOfferApiResponse) (int, int, error)
}

type OffersRepository interface {
	InsertOfferIfNotExists(id, summary string, contributors models.Contributors, image string) (bool, error)
	FindOffersByBankAndInstrument(ctx context.Context, bankName, paymentInstrument string) ([]models.Offer, error)
}
