package providers

import (
	"PERSONAL/ad_space_auction_service/services/adspace"
	"PERSONAL/ad_space_auction_service/services/auction"
	"PERSONAL/ad_space_auction_service/services/bidder"
)

var (
	AdspaceSvc adspace.AdspaceInterface = adspace.AdspaceImplementations{}
	BidderSvc  bidder.BidderInterface   = bidder.BidderImplementations{}
	AuctionSvc auction.AuctionInterface = auction.AuctionImplementation{}
)
