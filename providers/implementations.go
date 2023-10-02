package providers

import (
	"github.com/samsunil1999/ad_space_auction_service/services/adspace"
	"github.com/samsunil1999/ad_space_auction_service/services/auction"
	"github.com/samsunil1999/ad_space_auction_service/services/bidder"
)

var (
	AdspaceSvc adspace.AdspaceInterface = adspace.AdspaceImplementations{}
	BidderSvc  bidder.BidderInterface   = bidder.BidderImplementations{}
	AuctionSvc auction.AuctionInterface = auction.AuctionImplementation{}
)
