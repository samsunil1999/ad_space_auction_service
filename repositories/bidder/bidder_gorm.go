package bidder

import (
	"PERSONAL/ad_space_auction_service/database"
	"PERSONAL/ad_space_auction_service/models/entities"
)

type BidderRepoImpl struct{}

func (b BidderRepoImpl) GetById(id string) (entities.Bidders, error) {
	var record entities.Bidders
	err := database.Db.Where("uuid = ?", id).
		First(&record).Error

	return record, err
}

func (b BidderRepoImpl) Create(req entities.Bidders) (entities.Bidders, error) {
	err := database.Db.Create(&req).Error

	return req, err
}
