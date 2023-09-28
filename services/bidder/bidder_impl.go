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
		Uuid:        "bdr_" + uuid.NewString()[:23],
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
	bidders, err := repositories.BidderRepo.GetAll()
	if err != nil {
		return models.ListAllBiddersResp{}, err
	}
	return models.ListAllBiddersResp{
		Count: len(bidders),
		Data:  bidders,
	}, nil

}

func (b BidderImplementations) GetBidderById(id string) (entities.Bidders, error) {
	bidder, err := repositories.BidderRepo.GetById(id)
	return bidder, err
}

func (b BidderImplementations) UpdateBidderDetails(id string, req models.BidderReq) (entities.Bidders, error) {
	values := make(map[string]interface{})
	if req.Email != "" {
		values["email"] = req.Email
	}
	if req.Name != "" {
		values["name"] = req.Name
	}
	if req.PhoneNumber != "" {
		values["phone_number"] = req.PhoneNumber
	}

	return repositories.BidderRepo.UpdateWithCondition(id, values)
}

func (b BidderImplementations) DeleteBidder(id string) (models.DeleteBidderResp, error) {
	return models.DeleteBidderResp{}, nil
}
