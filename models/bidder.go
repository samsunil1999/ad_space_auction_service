package models

import (
	"PERSONAL/ad_space_auction_service/transformers"
)

type BidderReq struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
}

type ListAllBiddersResp struct {
	Count int `json:"count"`
	Data  []transformers.Bidders
}

type DeleteBidderResp struct {
	Message string `json:"message"`
}
