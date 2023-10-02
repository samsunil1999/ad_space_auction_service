package models

import "github.com/samsunil1999/ad_space_auction_service/transformers"

type AuctionResp struct {
	Adspace    transformers.Adspace `json:"adspace"`
	ActiveBids []transformers.Bids  `json:"active_bids,omitempty"`
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
