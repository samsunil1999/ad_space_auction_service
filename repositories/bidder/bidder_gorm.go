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

func (b BidderRepoImpl) GetAll() ([]entities.Bidders, error) {
	var bidders []entities.Bidders
	err := database.Db.Find(&bidders).Error

	return bidders, err
}

func (b BidderRepoImpl) UpdateWithCondition(id string, values map[string]interface{}) (entities.Bidders, error) {
	var record, updatedRecord entities.Bidders

	err := database.Db.Model(&record).Where("uuid = ?", id).
		Updates(values).Take(&updatedRecord).Error
	if err != nil {
		return entities.Bidders{}, err
	}

	return updatedRecord, nil
}
