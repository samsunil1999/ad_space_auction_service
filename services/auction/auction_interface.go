package auction

import (
	"PERSONAL/ad_space_auction_service/models"
	"PERSONAL/ad_space_auction_service/models/entities"
)

type AuctionInterface interface {
	GetAllLiveAuctions() (models.ListAuctionResp, error)
	GetAuctionById(id string) (models.AuctionResp, error)
	NewBidOnAuction(req models.BidOnAuctionReq) (entities.Bids, error)
}
