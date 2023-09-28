package bid

import "PERSONAL/ad_space_auction_service/models/entities"

type BidRepoInterface interface {
	GetAllByAdspaceId(id string) ([]entities.Bids, error)
	DeleteAllByAdspaceId(id string) error
	Create(entities.Bids) (entities.Bids, error)
}
