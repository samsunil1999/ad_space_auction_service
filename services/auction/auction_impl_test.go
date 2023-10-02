package auction

import (
	"PERSONAL/ad_space_auction_service/configs"
	"PERSONAL/ad_space_auction_service/database"
	"PERSONAL/ad_space_auction_service/models"
	"PERSONAL/ad_space_auction_service/models/entities"
	"PERSONAL/ad_space_auction_service/services/adspace"
	"PERSONAL/ad_space_auction_service/services/bidder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	configs.LoadConfig()
	database.Connect()
	database.SyncDatabase()
}

func TestGetAllLiveAuctions(t *testing.T) {
	req := models.AdspaceReq{
		Name:           "sam sunil",
		Description:    "test decription",
		BasePrice:      10.05,
		AuctionEndTime: "02-10-2023 19:05:00",
		ExpiredAt:      "02-10-2023 19:10:00",
	}
	createRes, err := adspace.AdspaceImplementations{}.CreateAdspace(req)
	assert.Nil(t, err)

	resp, err := AuctionImplementation{}.GetAllLiveAuctions()
	assert.Nil(t, err)
	found := false
	for _, val := range resp.Data {
		if val.Adspace.Uuid == createRes.Uuid {
			found = true
			break
		}
	}
	assert.True(t, found)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", createRes.Uuid).Delete(&entities.AdSpaces{})
	})
}

func TestGetAuctionById(t *testing.T) {
	req := models.AdspaceReq{
		Name:           "sam sunil",
		Description:    "test decription",
		BasePrice:      10.05,
		AuctionEndTime: "02-10-2023 19:05:00",
		ExpiredAt:      "02-10-2023 19:10:00",
	}
	createRes, err := adspace.AdspaceImplementations{}.CreateAdspace(req)
	assert.Nil(t, err)

	resp, err := AuctionImplementation{}.GetAuctionById(createRes.Uuid)
	assert.Nil(t, err)
	assert.Equal(t, createRes.Uuid, resp.Adspace.Uuid)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", createRes.Uuid).Delete(&entities.AdSpaces{})
	})
}

func TestNewBidOnAuction(t *testing.T) {
	adspacerReq := models.AdspaceReq{
		Name:           "sam sunil",
		Description:    "test decription",
		BasePrice:      10.05,
		AuctionEndTime: "02-10-2023 19:05:00",
		ExpiredAt:      "02-10-2023 19:10:00",
	}
	adspaceRes, err := adspace.AdspaceImplementations{}.CreateAdspace(adspacerReq)
	assert.Nil(t, err)

	bidderReq := models.BidderReq{
		Name:        "sam sunil",
		Email:       "xyz@gmail.com",
		PhoneNumber: "1234567890",
	}
	bidderRes, err := bidder.BidderImplementations{}.RegisterBidder(bidderReq)
	assert.Nil(t, err)

	req := models.BidOnAuctionReq{
		AdspaceId: adspaceRes.Uuid,
		BidderId:  bidderRes.Uuid,
		Amount:    adspaceRes.BasePrice + 10.5,
	}

	resp, err := AuctionImplementation{}.NewBidOnAuction(req)
	assert.Nil(t, err)
	assert.Equal(t, adspaceRes.Uuid, resp.AdSpaceId)
	assert.Equal(t, bidderRes.Uuid, resp.BidderId)
	assert.Equal(t, req.Amount, resp.BidAmount)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", resp.Uuid).Delete(&entities.Bids{})
		database.Db.Where("uuid = ?", adspaceRes.Uuid).Delete(&entities.AdSpaces{})
		database.Db.Where("uuid = ?", bidderRes.Uuid).Delete(&entities.Bidders{})
	})
}
