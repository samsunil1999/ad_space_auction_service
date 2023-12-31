package adspace

import "github.com/samsunil1999/ad_space_auction_service/models/entities"

type AdspaceRepoInterface interface {
	Create(req entities.AdSpaces) (entities.AdSpaces, error)
	GetAll() ([]entities.AdSpaces, error)
	GetById(id string) (entities.AdSpaces, error)
	UpdateWithCondition(id string, values map[string]interface{}) (entities.AdSpaces, error)
	DeleteById(id string) error
	GetAllWithStatus(status string) ([]entities.AdSpaces, error)
}
