package bidder

import (
	"github.com/samsunil1999/ad_space_auction_service/models"
	"github.com/samsunil1999/ad_space_auction_service/models/entities"
)

type BidderInterface interface {
	RegisterBidder(req models.BidderReq) (entities.Bidders, error)
	GetAllBidders() (models.ListAllBiddersResp, error)
	GetBidderById(id string) (entities.Bidders, error)
	UpdateBidderDetails(id string, req models.BidderReq) (entities.Bidders, error)
	DeleteBidder(id string) (models.DeleteBidderResp, error)
}
