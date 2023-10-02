package adspace

import (
	"testing"

	"github.com/samsunil1999/ad_space_auction_service/configs"
	"github.com/samsunil1999/ad_space_auction_service/constants"
	"github.com/samsunil1999/ad_space_auction_service/database"
	"github.com/samsunil1999/ad_space_auction_service/models"
	"github.com/samsunil1999/ad_space_auction_service/models/entities"

	"github.com/stretchr/testify/assert"
)

func init() {
	configs.LoadConfig()
	database.Connect()
	database.SyncDatabase()
}

func TestCreateAdspace(t *testing.T) {
	req := models.AdspaceReq{
		Name:           "sam sunil",
		Description:    "test decription",
		BasePrice:      10.05,
		AuctionEndTime: "02-10-2023 19:05:00",
		ExpiredAt:      "02-10-2023 19:10:00",
	}
	res, err := AdspaceImplementations{}.CreateAdspace(req)
	assert.Nil(t, err)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.Description, res.Description)
	assert.Equal(t, req.BasePrice, res.BasePrice)
	assert.Equal(t, req.AuctionEndTime, res.AuctionEndTime.Format(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS))
	assert.Equal(t, req.ExpiredAt, res.ExpiredAt.Format(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS))
	t.Cleanup(func() {
		database.Db.Where("uuid = ?", res.Uuid).Delete(&entities.AdSpaces{})
	})
}

func TestGetAllAvailableAdspace(t *testing.T) {
	// creating an adspace before listing
	req := models.AdspaceReq{
		Name:           "sam sunil",
		Description:    "test decription",
		BasePrice:      10.05,
		AuctionEndTime: "02-10-2023 19:05:00",
		ExpiredAt:      "02-10-2023 19:10:00",
	}
	createRes, err := AdspaceImplementations{}.CreateAdspace(req)
	assert.Nil(t, err)

	resp, err := AdspaceImplementations{}.GetAllAvailableAdspace()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, resp.Count, 1)
	found := false
	for _, adspace := range resp.Data {
		if adspace.Uuid == createRes.Uuid {
			found = true
			break
		}
	}
	assert.Equal(t, found, true)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", createRes.Uuid).Delete(&entities.AdSpaces{})
	})
}

func TestGetAdspaceById(t *testing.T) {
	// Creating an adspace before fetching by its uuid
	req := models.AdspaceReq{
		Name:           "sam sunil",
		Description:    "test decription",
		BasePrice:      10.05,
		AuctionEndTime: "02-10-2023 19:05:00",
		ExpiredAt:      "02-10-2023 19:10:00",
	}
	createRes, err := AdspaceImplementations{}.CreateAdspace(req)
	assert.Nil(t, err)

	resp, err := AdspaceImplementations{}.GetAdspaceById(createRes.Uuid)
	assert.Nil(t, err)
	assert.Equal(t, resp.Name, createRes.Name)
	assert.Equal(t, resp.Description, createRes.Description)
	assert.Equal(t, resp.BasePrice, createRes.BasePrice)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", createRes.Uuid).Delete(&entities.AdSpaces{})
	})

}

func TestUpdateAdspaceById(t *testing.T) {
	// Creating an adspace before updating by its uuid
	req := models.AdspaceReq{
		Name:           "sam sunil",
		Description:    "test decription",
		BasePrice:      10.05,
		AuctionEndTime: "02-10-2023 19:05:00",
		ExpiredAt:      "02-10-2023 19:10:00",
	}
	createRes, err := AdspaceImplementations{}.CreateAdspace(req)
	assert.Nil(t, err)

	resp, err := AdspaceImplementations{}.UpdateAdspaceById(createRes.Uuid, models.AdspaceReq{
		Name:        "updated name",
		Description: "updated description",
	})
	assert.Nil(t, err)
	assert.Equal(t, "updated name", resp.Name)
	assert.Equal(t, "updated description", resp.Description)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", createRes.Uuid).Delete(&entities.AdSpaces{})
	})
}

func TestDeleteAdspaceById(t *testing.T) {
	req := models.AdspaceReq{
		Name:           "sam sunil",
		Description:    "test decription",
		BasePrice:      10.05,
		AuctionEndTime: "02-10-2023 19:05:00",
		ExpiredAt:      "02-10-2023 19:10:00",
	}
	createRes, err := AdspaceImplementations{}.CreateAdspace(req)
	assert.Nil(t, err)

	resp, err := AdspaceImplementations{}.DeleteAdspaceById(createRes.Uuid)
	assert.Nil(t, err)
	assert.Contains(t, resp.Message, createRes.Uuid)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", createRes.Uuid).Delete(&entities.AdSpaces{})
	})
}
