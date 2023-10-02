package validators

import (
	"errors"

	"github.com/samsunil1999/ad_space_auction_service/models"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func ValidateRegisterBidderReq(c *gin.Context) (req models.BidderReq, err error) {
	err = c.ShouldBindJSON(&req)
	if err != nil {
		return models.BidderReq{}, err
	}

	opts := govalidator.Options{
		Data:  &req,
		Rules: getRulesForRegisterBidderReq(),
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		for param, message := range e {
			return req, errors.New("param: " + param + ", message:" + message[0])
		}
	}

	return req, nil
}

func getRulesForRegisterBidderReq() govalidator.MapData {
	rules := govalidator.MapData{
		"name":  []string{"required"},
		"email": []string{"required", "email"},
		"phone": []string{"required", "digits:10"},
	}
	return rules
}

func ValidateBidderId(c *gin.Context) (string, error) {
	id := c.Param("id")
	if id[:4] != "bdr_" {
		return id, errors.New("invalid id " + id)
	}

	return id, nil
}

func ValidateUpdateBidderReq(c *gin.Context) (req models.BidderReq, id string, err error) {
	id, err = ValidateBidderId(c)
	if err != nil {
		return models.BidderReq{}, "", err
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		return models.BidderReq{}, "", err
	}

	opts := govalidator.Options{
		Data:  &req,
		Rules: getRulesForUpdateBidderReq(),
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		for param, message := range e {
			return req, "", errors.New("param: " + param + ", message:" + message[0])
		}
	}

	return req, id, nil
}

func getRulesForUpdateBidderReq() govalidator.MapData {
	rules := govalidator.MapData{
		"email": []string{"email"},
		"phone": []string{"digits:10"},
	}
	return rules
}
