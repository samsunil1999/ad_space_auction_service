package adspace

import (
	"PERSONAL/ad_space_auction_service/constants"
	"PERSONAL/ad_space_auction_service/database"
	"PERSONAL/ad_space_auction_service/models"
	"PERSONAL/ad_space_auction_service/models/entities"
	"PERSONAL/ad_space_auction_service/providers/repositories"
	"errors"
	"log"

	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdspaceImplementations struct{}

func (a AdspaceImplementations) CreateAdspace(req models.AdspaceReq) (entities.AdSpaces, error) {
	roudedBasePrice := fmt.Sprintf("%.2f", req.BasePrice)
	basePrice, _ := strconv.ParseFloat(roudedBasePrice, 32)
	auctionEndTime, _ := time.Parse(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS, req.AuctionEndTime)
	expiredAt, _ := time.Parse(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS, req.ExpiredAt)

	adspace, err := repositories.AdspaceRepo.Create(entities.AdSpaces{
		Uuid:           "ads_" + uuid.NewString()[:23],
		Name:           req.Name,
		Description:    req.Description,
		BasePrice:      float32(basePrice),
		AuctionEndTime: auctionEndTime,
		ExpiredAt:      expiredAt,
		Status:         constants.Adspace_Status_IN_AUCTION,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	if err != nil {
		return entities.AdSpaces{}, err
	}

	// schedule a go routine to update adspace with bidder details
	go func() {
		auctionEndDuration := time.Since(auctionEndTime)

		time.AfterFunc(auctionEndDuration, func() {
			adspaceAfterAuctionEnd(adspace.Uuid)
		})
	}()

	//  schedule a go routine to delete adspace after expiry
	go func() {
		expiredAtDuration := time.Since(expiredAt)
		time.AfterFunc(expiredAtDuration, func() {
			adspaceAfterExpiredAt(adspace.Uuid)
		})
	}()
	return adspace, nil
}

func (a AdspaceImplementations) GetAllAvailableAdspace() (models.ListAllAdspaceResp, error) {
	data, err := repositories.AdspaceRepo.GetAll()
	if err != nil {
		return models.ListAllAdspaceResp{}, err
	}

	return models.ListAllAdspaceResp{
		Count: len(data),
		Data:  data,
	}, nil
}

func (a AdspaceImplementations) GetAdspaceById(id string) (entities.AdSpaces, error) {
	adspace, err := repositories.AdspaceRepo.GetById(id)
	if err != nil {
		return entities.AdSpaces{}, err
	}

	return adspace, nil
}

func (a AdspaceImplementations) UpdateAdspaceById(id string, req models.AdspaceReq) (entities.AdSpaces, error) {
	bids, err := repositories.BidRepo.GetAllByAdspaceId(id)
	if err != nil {
		return entities.AdSpaces{}, err
	}

	if len(bids) > 0 && req.BasePrice != 0.0 {
		return entities.AdSpaces{}, errors.New("Cannot update base_price of adspace with active bids")
	}

	values := make(map[string]interface{})
	if req.Name != "" {
		values["name"] = req.Name
	}
	if req.BasePrice > 0.0 {
		values["base_price"] = req.BasePrice
	}
	if req.Description != "" {
		values["description"] = req.Description
	}

	return repositories.AdspaceRepo.UpdateWithCondition(id, values)
}

func (a AdspaceImplementations) DeleteAdspaceById(id string) (models.DeleteAdspaceResp, error) {
	bids, err := repositories.BidRepo.GetAllByAdspaceId(id)
	if err != nil {
		return models.DeleteAdspaceResp{}, err
	}

	err = database.Db.Transaction(func(tx *gorm.DB) error {
		// delete bids if available
		if len(bids) > 0 {
			err := repositories.BidRepo.DeleteAllByAdspaceId(id)
			if err != nil {
				return err
			}
		}

		err := repositories.AdspaceRepo.DeleteById(id)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return models.DeleteAdspaceResp{}, err
	}

	return models.DeleteAdspaceResp{
		Message: "adspace with id " + id + " deleted successfully",
	}, nil
}

func adspaceAfterExpiredAt(id string) {
	// delete adspace
	err := repositories.AdspaceRepo.DeleteById(id)
	if err != nil {
		log.Println("[adspaceAfterExpiredAt][AdspaceRepo][DeleteById] failed::", id, "::error::", err.Error())
		return
	}
}

func adspaceAfterAuctionEnd(id string) {
	// fetch from adspace
	adspace, err := repositories.AdspaceRepo.GetById(id)
	if err != nil {
		log.Println("[adspaceAfterAuctionEnd][AdspaceRepo][GetById] failed::", adspace.Uuid, "::error::", err.Error())
		return
	}

	// fetch from bidders for highest bid
	bids, err := repositories.BidRepo.GetAllByAdspaceId(adspace.Uuid)
	if err != nil {
		log.Println("[adspaceAfterAuctionEnd][BidRepo][GetAllByAdspaceId] failed::", adspace.Uuid, "::error::", err.Error())
		return
	}

	// if there is no bids then update the status to 'expired'
	if len(bids) == 0 {
		log.Println("[adspaceAfterAuctionEnd] no bids found for adspace::", adspace.Uuid)

		_, err = repositories.AdspaceRepo.UpdateWithCondition(adspace.Uuid, map[string]interface{}{"status": constants.Adspace_Status_EXPIRED})
		if err != nil {
			log.Println("[adspaceAfterAuctionEnd][AdspaceRepo][UpdateWithCondition] failed::", adspace.Uuid, "::error::", err.Error())
		}
		return
	}

	// find the max bid
	maxBidAmount := adspace.BasePrice
	maxBidderId := ""
	for _, bid := range bids {
		if maxBidAmount < bid.BidAmount {
			maxBidAmount = bid.BidAmount
			maxBidderId = bid.BidderId
		}
	}

	// Update adspace with status 'live' and max bidder details
	_, err = repositories.AdspaceRepo.UpdateWithCondition(adspace.Uuid, map[string]interface{}{
		"status":     constants.Adspce_Status_LIVE,
		"sold_price": maxBidAmount,
		"bidder_id":  maxBidderId,
	})
	if err != nil {
		log.Println("[adspaceAfterAuctionEnd][AdspaceRepo][UpdateWithCondition] failed::", adspace.Uuid, "::error::", err.Error())
		return
	}
}
