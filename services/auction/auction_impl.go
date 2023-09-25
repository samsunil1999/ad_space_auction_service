package auction

import "PERSONAL/ad_space_auction_service/models"

type AuctionImplementation struct{}

func (a AuctionImplementation) GetAllLiveAuctions() (models.ListAuctionResp, error) {
	return models.ListAuctionResp{}, nil
}

func (a AuctionImplementation) GetAuctionById(id string) (models.AuctionResp, error) {
	return models.AuctionResp{}, nil
}

func (a AuctionImplementation) NewBidOnAuction(req models.BidOnAuctionReq) (models.BidOnAuctionResp, error) {
	return models.BidOnAuctionResp{}, nil
}
