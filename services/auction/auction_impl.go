package auction

import (
	"PERSONAL/ad_space_auction_service/constants"
	"PERSONAL/ad_space_auction_service/models"
	"PERSONAL/ad_space_auction_service/models/entities"
	"PERSONAL/ad_space_auction_service/providers/repositories"
	"errors"
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
	adspace, err := repositories.AdspaceRepo.GetById(id)
	if err != nil {
		return models.AuctionResp{}, err
	}

	// get all active bids if the adspace is in in_auction state
	var bids []entities.Bids
	if adspace.Status == constants.Adspace_Status_IN_AUCTION {
		bids, err = repositories.BidRepo.GetAllByAdspaceId(id)
		if err != nil {
			return models.AuctionResp{}, err
		}
	} else {
		return models.AuctionResp{}, errors.New("adspace " + id + " is not in auction")
	}

	return models.AuctionResp{
		Adspace:    adspace,
		ActiveBids: bids,
	}, nil
}

func (a AuctionImplementation) NewBidOnAuction(req models.BidOnAuctionReq) (models.BidOnAuctionResp, error) {
	return models.BidOnAuctionResp{}, nil
}
