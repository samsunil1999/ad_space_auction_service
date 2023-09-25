package adspace

import (
	"PERSONAL/ad_space_auction_service/database"
	"PERSONAL/ad_space_auction_service/models/entities"
)

type AdspaceRepoImpl struct{}

func (a AdspaceRepoImpl) Create(record entities.AdSpaces) (entities.AdSpaces, error) {

	err := database.Db.Create(&record).Error
	if err != nil {
		return entities.AdSpaces{}, err
	}

	return record, nil
}

func (a AdspaceRepoImpl) GetAll() ([]entities.AdSpaces, error) {
	var records []entities.AdSpaces
	err := database.Db.Find(&records).Error
	if err != nil {
		return []entities.AdSpaces{}, err
	}

	return records, nil
}

func (a AdspaceRepoImpl) GetById(id string) (entities.AdSpaces, error) {
	var record entities.AdSpaces
	err := database.Db.Where("uuid = ?", id).
		First(&record).Error
	if err != nil {
		return entities.AdSpaces{}, err
	}

	return record, nil
}

func (a AdspaceRepoImpl) UpdateWithCondition(id string, values map[string]interface{}) (entities.AdSpaces, error) {
	var record entities.AdSpaces
	var updatedRecord entities.AdSpaces
	err := database.Db.Model(&record).Where("uuid = ?", id).
		Updates(values).Take(&updatedRecord).Error
	if err != nil {
		return entities.AdSpaces{}, err
	}

	return updatedRecord, nil
}

func (a AdspaceRepoImpl) DeleteById(id string) (err error) {
	err = database.Db.Where("uuid = ?", id).Delete(&entities.AdSpaces{}).Error
	return err
}