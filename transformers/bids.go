package transformers

import (
	"PERSONAL/ad_space_auction_service/constants"
	"PERSONAL/ad_space_auction_service/models/entities"
	"time"
)

type Bids struct {
	ID        int64   `json:"ID"`
	Uuid      string  `json:"uuid"`
	AdSpaceId string  `json:"ad_space_id"`
	BidderId  string  `json:"bidder_id"`
	BidAmount float32 `json:"bid_amount"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt string `json:"deleted_at,omitempty"`
}

func GetBidsModel(req entities.Bids) Bids {
	var deletedAt string
	if req.DeletedAt.Valid {
		deletedAt = req.DeletedAt.Time.Format(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS)
	}

	return Bids{
		ID:        req.ID,
		Uuid:      req.Uuid,
		AdSpaceId: req.AdSpaceId,
		BidderId:  req.BidderId,
		BidAmount: req.BidAmount,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
		DeletedAt: deletedAt,
	}
}
