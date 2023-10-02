package bid

import "github.com/samsunil1999/ad_space_auction_service/models/entities"

type BidRepoInterface interface {
	GetAllByAdspaceId(id string) ([]entities.Bids, error)
	Create(entities.Bids) (entities.Bids, error)
}
