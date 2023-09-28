package auction

import (
	"PERSONAL/ad_space_auction_service/constants"
	"PERSONAL/ad_space_auction_service/models"
	"PERSONAL/ad_space_auction_service/providers/repositories"
)

type AuctionImplementation struct{}

func (a AuctionImplementation) GetAllLiveAuctions() (models.ListAuctionResp, error) {
	adspaces, err := repositories.AdspaceRepo.GetAllWithStatus(constants.Adspace_Status_IN_AUCTION)
	if err != nil {
		return models.ListAuctionResp{}, err
	}

	var auctionData []models.AuctionResp
	for _, adspace := range adspaces {
		auctionData = append(auctionData, models.AuctionResp{
			Adspace: adspace,
		})
	}

	return models.ListAuctionResp{
		Count: len(adspaces),
		Data:  auctionData,
	}, nil
}

func (a AuctionImplementation) GetAuctionById(id string) (models.AuctionResp, error) {
	return models.AuctionResp{}, nil
}

func (a AuctionImplementation) NewBidOnAuction(req models.BidOnAuctionReq) (models.BidOnAuctionResp, error) {
	return models.BidOnAuctionResp{}, nil
}
