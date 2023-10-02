package repositories

import (
	"github.com/samsunil1999/ad_space_auction_service/repositories/adspace"
	"github.com/samsunil1999/ad_space_auction_service/repositories/bid"
	"github.com/samsunil1999/ad_space_auction_service/repositories/bidder"
)

var (
	AdspaceRepo adspace.AdspaceRepoInterface = adspace.AdspaceRepoImpl{}
	BidRepo     bid.BidRepoInterface         = bid.BidRepoImpl{}
	BidderRepo  bidder.BidderRepoInterface   = bidder.BidderRepoImpl{}
)
