package controllers

import (
	"awesomeProject2/interfaces"
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OfferController struct {
	Service interfaces.OffersService
}

func NewOfferController(service interfaces.OffersService) *OfferController {
	return &OfferController{Service: service}
}

type OfferRequest struct {
	FlipkartOfferApiResponse models.FlipkartOfferApiResponse `json:"flipkartOfferApiResponse"`
}

func (c *OfferController) PostOffer(ctx *gin.Context) {
	var req OfferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total, created, err := c.Service.SaveOffers(req.FlipkartOfferApiResponse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"noOfOffersIdentified": total,
		"noOfNewOffersCreated": created,
	})
}
