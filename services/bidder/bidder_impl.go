package bidder

import (
	"PERSONAL/ad_space_auction_service/models"
	"PERSONAL/ad_space_auction_service/models/entities"
)

type BidderImplementations struct{}

func (b BidderImplementations) RegisterBidder(req models.BidderReq) (entities.Bidders, error) {
	return entities.Bidders{}, nil
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
