package bidder

import "PERSONAL/ad_space_auction_service/models/entities"

type BidderRepoInterface interface {
	GetById(id string) (entities.Bidders, error)
	Create(entities.Bidders) (entities.Bidders, error)
	GetAll() ([]entities.Bidders, error)
}
