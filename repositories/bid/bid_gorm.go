package bid

import (
	"PERSONAL/ad_space_auction_service/database"
	"PERSONAL/ad_space_auction_service/models/entities"
)

type BidRepoImpl struct{}

func (b BidRepoImpl) GetAllByAdspaceId(id string) (bids []entities.Bids, err error) {
	err = database.Db.Where("ad_space_id = ? AND deleted_at IS NULL", id).
		Find(&bids).Error
	if err != nil {
		return []entities.Bids{}, err
	}

	return bids, nil
}

func (b BidRepoImpl) Create(bid entities.Bids) (entities.Bids, error) {
	err := database.Db.Create(&bid).Error
	if err != nil {
		return entities.Bids{}, err
	}

	return bid, nil
}
