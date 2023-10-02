package validators

import (
	"PERSONAL/ad_space_auction_service/constants"
	"PERSONAL/ad_space_auction_service/models"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func ValiateCreateAdspaceReq(ctx *gin.Context) (req models.AdspaceReq, err error) {
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		return req, err
	}

	opts := govalidator.Options{
		Data:  &req,
		Rules: getRulesForCreateAdspaceReq(),
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		for param, message := range e {
			return req, errors.New("param: " + param + ", message:" + message[0])
		}
	}

	auctionEndTime, err := time.Parse(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS, req.AuctionEndTime)
	if err != nil {
		return models.AdspaceReq{}, errors.New("invalid auction_end_time, must be in format dd-mm-yyyy hh:mm:ss")
	}

	expiredAt, err := time.Parse(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS, req.ExpiredAt)
	if err != nil {
		return models.AdspaceReq{}, errors.New("invalid  expired_at, must be in format dd-mm-yyyy hh:mm:ss")
	}

	if auctionEndTime.Before(time.Now().Add((5 * time.Hour) + (30 * time.Minute))) {
		return models.AdspaceReq{}, errors.New("auction end time should be a future time")
	}

	if expiredAt.Before(auctionEndTime) {
		return models.AdspaceReq{}, errors.New("expired_at cannot be less than auction_end_time")
	}

	if req.BasePrice < 0.0 {
		return models.AdspaceReq{}, errors.New("Base price cannot be negative")
	}

	return req, nil
}

func getRulesForCreateAdspaceReq() govalidator.MapData {
	rules := govalidator.MapData{
		"name":             []string{"required"},
		"description":      []string{"required"},
		"base_price":       []string{"required", "float"},
		"auction_end_time": []string{"required"},
		"expired_at":       []string{"required"},
	}
	return rules
}

func ValidateGetAdspaceById(ctx *gin.Context) (id string, err error) {
	id = ctx.Param("id")
	if id[:4] != "ads_" {
		return id, errors.New("invalid id " + id)
	}

	return id, nil
}

func ValidateUpdateAdspaceByIdReq(ctx *gin.Context) (req models.AdspaceReq, id string, err error) {

	id, err = ValidateGetAdspaceById(ctx)
	if err != nil {
		return models.AdspaceReq{}, "", err
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		return models.AdspaceReq{}, "", err
	}

	opts := govalidator.Options{
		Data:  &req,
		Rules: getRulesForUpdateAdspaceReq(),
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		for param, message := range e {
			return req, "", errors.New("param: " + param + ", message:" + message[0])
		}
	}

	if req.AuctionEndTime != "" {
		return models.AdspaceReq{}, "", errors.New("auction_end_time cannot be updated")
	}

	if req.ExpiredAt != "" {
		return models.AdspaceReq{}, "", errors.New("expired_at cannot be less before auction_end_time")
	}

	if req.BasePrice < 0 {
		return models.AdspaceReq{}, "", errors.New("Base price cannot be less than 0")
	}

	return req, id, nil
}

func getRulesForUpdateAdspaceReq() govalidator.MapData {
	rules := govalidator.MapData{
		"base_price": []string{"float"},
	}
	return rules
}
