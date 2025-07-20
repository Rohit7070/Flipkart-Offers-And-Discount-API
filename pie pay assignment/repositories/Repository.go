package repositories

import (
	"awesomeProject2/config"
	"awesomeProject2/interfaces"
	"awesomeProject2/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type offerRepository struct {
	collection *mongo.Collection
}

func NewOfferRepository() interfaces.OffersRepository {
	return &offerRepository{
		collection: config.DB.Collection("offers"),
	}
}

func (r *offerRepository) InsertOfferIfNotExists(id, summary string, contributors models.Contributors, image string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if exists
	count, err := r.collection.CountDocuments(ctx, bson.M{"adjustment_id": id})
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, nil // already exists
	}

	// Insert new
	_, err = r.collection.InsertOne(ctx, bson.M{
		"adjustment_id": id,
		"summary":       summary,
		"banks":         contributors.Banks,
		"payment_modes": contributors.PaymentInstrument,
		"emi_months":    contributors.EmiMonths,
		"image":         image,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *offerRepository) FindOffersByBankAndInstrument(ctx context.Context, bankName, paymentInstrument string) ([]models.Offer, error) {
	collection := config.DB.Collection("offers")
	var offers []models.Offer

	filter := bson.M{
		"banks": bson.M{"$in": []string{bankName}},
	}

	if paymentInstrument != "" {
		filter["payment_modes"] = bson.M{"$in": []string{paymentInstrument}}
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &offers); err != nil {
		return nil, err
	}
	return offers, nil
}
