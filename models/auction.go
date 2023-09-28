package models

import "PERSONAL/ad_space_auction_service/models/entities"

type AuctionResp struct {
	// Name           string  `json:"name"`
	// Description    string  `json:"description"`
	// BasePrice      float32 `json:"base_price"`
	// SoldPrice      float32 `json:"sold_price,omitempty"`
	// Status         string  `json:"status"`
	// AuctionEndTime string  `json:"auction_end_time"`
	// ExpiredAt      string  `json:"expired_at"`
	// CreatedAt      string  `json:"created_at"`
	// UpdatedAt      string  `json:"updated_at"`
	// DeletedAt      string  `json:"deleted_at,omitempty"`

	Adspace    entities.AdSpaces `json:"adspace"`
	ActiveBids []entities.Bids   `json:"active_bids,omitempty"`
}

type ListAuctionResp struct {
	Count int           `json:"count"`
	Data  []AuctionResp `json:"data"`
}

type BidOnAuctionReq struct {
	AdspaceId string  `json:"adspace_id"`
	BidderId  string  `json:"bidder_id"`
	Amount    float32 `json:"amount"`
}
