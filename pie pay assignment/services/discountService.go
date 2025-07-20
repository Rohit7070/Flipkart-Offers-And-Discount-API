package services

import (
	"awesomeProject2/interfaces"
	"context"
	"regexp"
	"strconv"
	"strings"
)

type DiscountService struct {
	repo interfaces.OffersRepository
}

func NewDiscountService(repo interfaces.OffersRepository) *DiscountService {
	return &DiscountService{repo: repo}
}

func (s *DiscountService) GetHighestDiscount(ctx context.Context, bankName, paymentInstrument string) (float64, error) {
	offers, err := s.repo.FindOffersByBankAndInstrument(ctx, bankName, paymentInstrument)
	if err != nil {
		return 0, err
	}

	if len(offers) == 0 {
		return 0, nil
	}
	highest := 0.0
	for _, o := range offers {
		discount := ExtractDiscountAmount(o.Summary)
		if discount > highest {
			highest = discount
		}
	}
	return highest, nil
}

func ExtractDiscountAmount(summary string) float64 {
	summary = strings.ReplaceAll(summary, ",", "")

	reAmount := regexp.MustCompile(`₹\s?(\d+\.?\d*)`)
	matches := reAmount.FindStringSubmatch(summary)
	if len(matches) > 1 {
		amount, _ := strconv.ParseFloat(matches[1], 64)
		return amount
	}
	rePercent := regexp.MustCompile(`(\d+\.?\d*)\s?%`)
	matches = rePercent.FindStringSubmatch(summary)
	if len(matches) > 1 {
		percent, _ := strconv.ParseFloat(matches[1], 64)
		return percent
	}

	return 0
}
