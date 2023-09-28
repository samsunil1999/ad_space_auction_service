package bidder

import (
	"PERSONAL/ad_space_auction_service/models"
	"PERSONAL/ad_space_auction_service/models/entities"
	"PERSONAL/ad_space_auction_service/providers/repositories"
	"time"

	"github.com/google/uuid"
)

type BidderImplementations struct{}

func (b BidderImplementations) RegisterBidder(req models.BidderReq) (entities.Bidders, error) {

	bidder, err := repositories.BidderRepo.Create(entities.Bidders{
		Uuid:        "bid_" + uuid.NewString()[:23],
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return entities.Bidders{}, err
	}

	return bidder, nil
}

func (b BidderImplementations) GetAllBidders() (models.ListAllBiddersResp, error) {
	return models.ListAllBiddersResp{}, nil
}

func (b BidderImplementations) GetBidderById(id string) (entities.Bidders, error) {
	return entities.Bidders{}, nil
}

func (b BidderImplementations) UpdateBidderDetails(req models.BidderReq) (entities.Bidders, error) {
	return entities.Bidders{}, nil
}

func (b BidderImplementations) DeleteBidder(id string) (models.DeleteBidderResp, error) {
	return models.DeleteBidderResp{}, nil
}
