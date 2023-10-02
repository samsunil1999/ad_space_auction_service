package adspace

import (
	"github.com/samsunil1999/ad_space_auction_service/models"
	"github.com/samsunil1999/ad_space_auction_service/models/entities"
)

type AdspaceInterface interface {
	CreateAdspace(req models.AdspaceReq) (entities.AdSpaces, error)
	GetAllAvailableAdspace() (models.ListAllAdspaceResp, error)
	GetAdspaceById(id string) (entities.AdSpaces, error)
	UpdateAdspaceById(id string, req models.AdspaceReq) (entities.AdSpaces, error)
	DeleteAdspaceById(id string) (models.DeleteAdspaceResp, error)
}
