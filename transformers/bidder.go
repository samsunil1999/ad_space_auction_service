package transformers

import (
	"PERSONAL/ad_space_auction_service/constants"
	"PERSONAL/ad_space_auction_service/models/entities"
	"time"
)

type Bidders struct {
	ID          int64  `json:"ID"`
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   string `json:"deleted_at,omitempty"`
}

func GetBidderModel(req entities.Bidders) Bidders {
	var deletedAt string
	if req.DeletedAt.Valid {
		deletedAt = req.DeletedAt.Time.Format(constants.Time_Format_DD_MM_YYYY_WITH_COLON_HH_MM_SS)
	}

	return Bidders{
		ID:          req.ID,
		Uuid:        req.Uuid,
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}
