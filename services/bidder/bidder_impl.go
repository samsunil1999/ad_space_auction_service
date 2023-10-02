package bidder

import (
	"time"

	"github.com/samsunil1999/ad_space_auction_service/constants"
	"github.com/samsunil1999/ad_space_auction_service/database"
	"github.com/samsunil1999/ad_space_auction_service/models"
	"github.com/samsunil1999/ad_space_auction_service/models/entities"
	"github.com/samsunil1999/ad_space_auction_service/providers/repositories"
	"github.com/samsunil1999/ad_space_auction_service/transformers"

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

	var bidderResp []transformers.Bidders
	for _, bidder := range bidders {
		bidderResp = append(bidderResp, transformers.GetBidderModel(bidder))
	}
	return models.ListAllBiddersResp{
		Count: len(bidders),
		Data:  bidderResp,
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
	tx := database.Db.Begin()
	if tx.Error != nil {
		return models.DeleteBidderResp{}, tx.Error
	}
	// Deleting all the bids made by bidder
	err := tx.Model(&entities.Bids{}).Where("bidder_id = ?", id).
		Update("DeletedAt", time.Now()).Error
	if err != nil {
		tx.Rollback()
		return models.DeleteBidderResp{}, err
	}

	// if there is an adspace with same bidder_id update status to expired and remove bidder_id
	err = tx.Model(&entities.AdSpaces{}).Where("bidder_id = ?", id).Updates(map[string]interface{}{
		"status":    constants.Adspace_Status_EXPIRED,
		"bidder_id": "",
	}).Error
	if err != nil {
		tx.Rollback()
		return models.DeleteBidderResp{}, err
	}

	// Delete the bidder
	err = tx.Model(&entities.Bidders{}).Where("uuid = ?", id).Update("DeletedAt", time.Now()).Error
	if err != nil {
		tx.Rollback()
		return models.DeleteBidderResp{}, err
	}

	if err = tx.Commit().Error; err != nil {
		return models.DeleteBidderResp{}, err
	}

	return models.DeleteBidderResp{
		Message: "Bidder " + id + "deleted successfuly ",
	}, nil
}
