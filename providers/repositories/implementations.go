package repositories

import (
	"PERSONAL/ad_space_auction_service/repositories/adspace"
	"PERSONAL/ad_space_auction_service/repositories/bid"
	"PERSONAL/ad_space_auction_service/repositories/bidder"
)

var (
	AdspaceRepo adspace.AdspaceRepoInterface = adspace.AdspaceRepoImpl{}
	BidRepo     bid.BidRepoInterface         = bid.BidRepoImpl{}
	BidderRepo  bidder.BidderRepoInterface   = bidder.BidderRepoImpl{}
)
