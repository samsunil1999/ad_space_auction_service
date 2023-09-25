package models

import "PERSONAL/ad_space_auction_service/models/entities"

type AdspaceReq struct {
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	BasePrice      float32 `json:"base_price"`
	AuctionEndTime string  `json:"auction_end_time"`
	ExpiredAt      string  `json:"expired_at"`
}

type ListAllAdspaceResp struct {
	Count int `json:"count"`
	Data  []entities.AdSpaces
}

type DeleteAdspaceResp struct {
	Message string `json:"message"`
}
