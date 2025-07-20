package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Offer struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AdjustmentID      string             `bson:"adjustment_id" json:"adjustment_id"`
	AdjustmentType    string             `bson:"adjustment_type,omitempty" json:"adjustment_type"`
	AdjustmentSubType string             `bson:"adjustment_sub_type,omitempty" json:"adjustment_sub_type"`
	Summary           string             `bson:"summary" json:"summary"`
	Banks             []string           `bson:"banks" json:"banks"`
	PaymentInstrument []string           `bson:"payment_instrument" json:"payment_instrument"`
	EmiMonths         []string           `bson:"emi_months" json:"emi_months"`
	DisplayTags       []string           `bson:"display_tags" json:"display_tags"`
	Image             string             `bson:"image" json:"image"`
	Type              string             `bson:"type,omitempty" json:"type"`
}

type FlipkartOfferApiResponse struct {
	OfferBanners  []OfferBanner           `json:"offer_banners"`
	OfferSections map[string]OfferSection `json:"offer_sections"`
}

type OfferBanner struct {
	AdjustmentSubType string       `json:"adjustment_sub_type"`
	AdjustmentID      string       `json:"adjustment_id"`
	Summary           string       `json:"summary"`
	Contributors      Contributors `json:"contributors"`
	DisplayTags       []string     `json:"display_tags"`
	Image             string       `json:"image"`
	Type              OfferType    `json:"type"`
}

type OfferSection struct {
	Title  string             `json:"title"`
	Offers []OfferSectionItem `json:"offers"`
}

type OfferSectionItem struct {
	AdjustmentType string       `json:"adjustment_type"`
	AdjustmentID   string       `json:"adjustment_id"`
	Summary        string       `json:"summary"`
	Contributors   Contributors `json:"contributors"`
	DisplayTags    []string     `json:"display_tags"`
	Image          string       `json:"image"`
}

type Contributors struct {
	PaymentInstrument []string `json:"payment_instrument"`
	Banks             []string `json:"banks"`
	EmiMonths         []string `json:"emi_months"`
	CardNetworks      []string `json:"card_networks"`
}

type OfferType struct {
	Value string `json:"value"`
}
