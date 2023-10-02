package auction

import (
	"PERSONAL/ad_space_auction_service/constants"
	"PERSONAL/ad_space_auction_service/models"
	"PERSONAL/ad_space_auction_service/models/entities"
	"PERSONAL/ad_space_auction_service/providers/repositories"
	"PERSONAL/ad_space_auction_service/transformers"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
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
			Adspace: transformers.GetAdspaceModel(adspace),
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
	var bidsResp []transformers.Bids
	if adspace.Status == constants.Adspace_Status_IN_AUCTION {
		bids, err := repositories.BidRepo.GetAllByAdspaceId(id)
		if err != nil {
			return models.AuctionResp{}, err
		}
		for _, bid := range bids {
			bidsResp = append(bidsResp, transformers.GetBidsModel(bid))
		}
	} else {
		return models.AuctionResp{}, errors.New("adspace " + id + " is not in auction")
	}

	return models.AuctionResp{
		Adspace:    transformers.GetAdspaceModel(adspace),
		ActiveBids: bidsResp,
	}, nil
}

func (a AuctionImplementation) NewBidOnAuction(req models.BidOnAuctionReq) (entities.Bids, error) {

	adspace, err := repositories.AdspaceRepo.GetById(req.AdspaceId)
	if err != nil {
		return entities.Bids{}, err
	}

	if req.Amount < adspace.BasePrice {
		return entities.Bids{}, errors.New("amount is less than base_price " + fmt.Sprintf("%2f", adspace.BasePrice))
	}

	// check if the bidder id exists
	_, err = repositories.BidderRepo.GetById(req.BidderId)
	if err != nil {
		return entities.Bids{}, err
	}

	bid, err := repositories.BidRepo.Create(entities.Bids{
		Uuid:      "bid_" + uuid.NewString()[:23],
		AdSpaceId: req.AdspaceId,
		BidderId:  req.BidderId,
		BidAmount: req.Amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return entities.Bids{}, err
	}

	return bid, nil
}
