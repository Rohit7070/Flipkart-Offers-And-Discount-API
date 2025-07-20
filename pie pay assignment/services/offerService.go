package services

import (
	"awesomeProject2/interfaces"
	"awesomeProject2/models"
)

type offerService struct {
	repo interfaces.OffersRepository
}

func NewOfferService(repo interfaces.OffersRepository) *offerService {
	return &offerService{repo: repo}
}

func (s *offerService) SaveOffers(apiResp models.FlipkartOfferApiResponse) (int, int, error) {
	totalOffers := 0
	newOffers := 0

	// Process Offer Banners
	for _, banner := range apiResp.OfferBanners {
		totalOffers++
		created, _ := s.repo.InsertOfferIfNotExists(banner.AdjustmentID, banner.Summary, banner.Contributors, banner.Image)
		if created {
			newOffers++
		}
	}

	// Process Offer Sections
	for _, section := range apiResp.OfferSections {
		for _, offer := range section.Offers {
			totalOffers++
			created, _ := s.repo.InsertOfferIfNotExists(offer.AdjustmentID, offer.Summary, offer.Contributors, offer.Image)
			if created {
				newOffers++
			}
		}
	}

	return totalOffers, newOffers, nil
}
