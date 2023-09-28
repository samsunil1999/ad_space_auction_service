package bidder

import "PERSONAL/ad_space_auction_service/models/entities"

type BidRepoInterface interface {
	GetById(id string) (entities.Bidders, error)
}
