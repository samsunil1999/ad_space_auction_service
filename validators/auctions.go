package validators

import (
	"PERSONAL/ad_space_auction_service/models"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func ValidateBidOnAuctionReq(c *gin.Context) (req models.BidOnAuctionReq, err error) {
	err = c.ShouldBindJSON(&req)
	if err != nil {
		return req, err
	}

	opts := govalidator.Options{
		Data:  &req,
		Rules: getRulesForBidOnAuctionReq(),
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		for param, message := range e {
			return req, errors.New("param: " + param + ", message:" + message[0])
		}
	}

	if req.AdspaceId[:4] != "ads_" {
		return models.BidOnAuctionReq{}, errors.New("Invalid adspace id " + req.AdspaceId)
	}

	if req.BidderId[:4] != "bdr_" {
		return models.BidOnAuctionReq{}, errors.New("Invalid bidder id " + req.BidderId)
	}

	return req, nil
}

func getRulesForBidOnAuctionReq() govalidator.MapData {
	rules := govalidator.MapData{
		"adspace_id": []string{"required"},
		"bidder_id":  []string{"required"},
		"amount":     []string{"required", "float"},
	}
	return rules
}
