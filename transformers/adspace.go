package transformers

import (
	"time"

	"github.com/samsunil1999/ad_space_auction_service/constants"
	"github.com/samsunil1999/ad_space_auction_service/models/entities"
)

type Adspace struct {
	ID             int64     `json:"ID"`
	Uuid           string    `json:"uuid"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	BasePrice      float32   `json:"base_price"`
	SoldPrice      float32   `json:"sold_price,omitempty"`
	BidderId       string    `json:"bidder_id,omitempty"`
	AuctionEndTime time.Time `json:"auction_end_time"`
	ExpiredAt      time.Time `json:"expired_at"`
	Status         string    `json:"status"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      string `json:"deleted_at,omitempty"`
}

func GetAdspaceModel(req entities.AdSpaces) Adspace {
	var deletedAt string
	if req.DeletedAt.Valid {
		deletedAt = req.DeletedAt.Time.Format(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS)
	}

	return Adspace{
		ID:             req.ID,
		Uuid:           req.Uuid,
		Name:           req.Name,
		Description:    req.Description,
		BasePrice:      req.BasePrice,
		SoldPrice:      req.SoldPrice,
		BidderId:       req.BidderId,
		AuctionEndTime: req.AuctionEndTime,
		ExpiredAt:      req.ExpiredAt,
		Status:         req.Status,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
		DeletedAt:      deletedAt,
	}
}
