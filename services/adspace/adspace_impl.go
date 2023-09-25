package adspace

import (
	"PERSONAL/ad_space_auction_service/constants"
	"PERSONAL/ad_space_auction_service/models"
	"PERSONAL/ad_space_auction_service/models/entities"
	"PERSONAL/ad_space_auction_service/providers/repositories"

	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type AdspaceImplementations struct{}

func (a AdspaceImplementations) CreateAdspace(req models.AdspaceReq) (entities.AdSpaces, error) {
	roudedBasePrice := fmt.Sprintf("%.2f", req.BasePrice)
	basePrice, _ := strconv.ParseFloat(roudedBasePrice, 32)
	auctionEndTime, _ := time.Parse(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS, req.AuctionEndTime)
	expiredAt, _ := time.Parse(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS, req.ExpiredAt)

	return repositories.AdspaceRepo.Create(entities.AdSpaces{
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
	if req.AuctionEndTime != "" {
		t, _ := time.Parse(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS, req.AuctionEndTime)
		values["auction_end_time"] = t
	}
	if req.ExpiredAt != "" {
		t, _ := time.Parse(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS, req.ExpiredAt)
		values["expired_at"] = t
	}

	return repositories.AdspaceRepo.UpdateWithCondition(id, values)
}

func (a AdspaceImplementations) DeleteAdspaceById(id string) (models.DeleteAdspaceResp, error) {
	err := repositories.AdspaceRepo.DeleteById(id)
	if err != nil {
		return models.DeleteAdspaceResp{}, err
	}

	return models.DeleteAdspaceResp{
		Message: "adspace with id " + id + " deleted successfully",
	}, nil
}
