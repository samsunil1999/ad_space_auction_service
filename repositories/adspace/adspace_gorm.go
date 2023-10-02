package adspace

import (
	"time"

	"github.com/samsunil1999/ad_space_auction_service/database"
	"github.com/samsunil1999/ad_space_auction_service/models/entities"
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
	err := database.Db.Where("uuid = ? AND deleted_at IS NULL", id).
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
	err = database.Db.Model(&entities.AdSpaces{}).Where("uuid = ?", id).Update("DeletedAt", time.Now()).Error
	return err
}

func (a AdspaceRepoImpl) GetAllWithStatus(status string) ([]entities.AdSpaces, error) {
	var records []entities.AdSpaces

	err := database.Db.Where("status = ? AND deleted_at IS NULL", status).
		Find(&records).Error
	if err != nil {
		return []entities.AdSpaces{}, err
	}

	return records, nil
}
