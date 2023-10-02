package bidder

import (
	"testing"

	"github.com/samsunil1999/ad_space_auction_service/configs"
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

func TestRegisterBidder(t *testing.T) {
	req := models.BidderReq{
		Name:        "sam sunil",
		Email:       "xyz@gmail.com",
		PhoneNumber: "1234567890",
	}
	resp, err := BidderImplementations{}.RegisterBidder(req)
	assert.Nil(t, err)
	assert.Equal(t, req.Name, resp.Name)
	assert.Equal(t, req.Email, resp.Email)
	assert.Equal(t, req.PhoneNumber, resp.PhoneNumber)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", resp.Uuid).Delete(&entities.Bidders{})
	})
}

func TestGetAllBidders(t *testing.T) {
	createReq := models.BidderReq{
		Name:        "sam sunil",
		Email:       "xyz@gmail.com",
		PhoneNumber: "1234567890",
	}
	createResp, err := BidderImplementations{}.RegisterBidder(createReq)
	assert.Nil(t, err)

	resp, err := BidderImplementations{}.GetAllBidders()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, resp.Count, 1)
	found := false
	for _, bdr := range resp.Data {
		if bdr.Uuid == createResp.Uuid {
			found = true
			break
		}
	}
	assert.True(t, found)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", createResp.Uuid).Delete(&entities.Bidders{})
	})
}

func TestGetBidderById(t *testing.T) {
	createReq := models.BidderReq{
		Name:        "sam sunil",
		Email:       "xyz@gmail.com",
		PhoneNumber: "1234567890",
	}
	createResp, err := BidderImplementations{}.RegisterBidder(createReq)
	assert.Nil(t, err)

	resp, err := BidderImplementations{}.GetBidderById(createResp.Uuid)
	assert.Nil(t, err)
	assert.Equal(t, createReq.Name, resp.Name)
	assert.Equal(t, createReq.Email, resp.Email)
	assert.Equal(t, createReq.PhoneNumber, resp.PhoneNumber)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", createResp.Uuid).Delete(&entities.Bidders{})
	})
}

func TestUpdateBidderDetails(t *testing.T) {
	createReq := models.BidderReq{
		Name:        "sam sunil",
		Email:       "xyz@gmail.com",
		PhoneNumber: "1234567890",
	}
	createResp, err := BidderImplementations{}.RegisterBidder(createReq)
	assert.Nil(t, err)

	req := models.BidderReq{
		Name:  "new name sam",
		Email: "new@gmail.com",
	}
	resp, err := BidderImplementations{}.UpdateBidderDetails(createResp.Uuid, req)
	assert.Nil(t, err)
	assert.Equal(t, req.Name, resp.Name)
	assert.Equal(t, req.Email, resp.Email)
	assert.Equal(t, createResp.PhoneNumber, resp.PhoneNumber)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", createResp.Uuid).Delete(&entities.Bidders{})
	})
}

func TestDeleteBidder(t *testing.T) {
	createReq := models.BidderReq{
		Name:        "sam sunil",
		Email:       "xyz@gmail.com",
		PhoneNumber: "1234567890",
	}
	createResp, err := BidderImplementations{}.RegisterBidder(createReq)
	assert.Nil(t, err)

	resp, err := BidderImplementations{}.DeleteBidder(createResp.Uuid)
	assert.Nil(t, err)
	assert.Contains(t, resp, createResp.Uuid)

	t.Cleanup(func() {
		database.Db.Where("uuid = ?", createResp.Uuid).Delete(&entities.Bidders{})
	})
}
