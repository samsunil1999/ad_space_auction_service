package repositories

import "PERSONAL/ad_space_auction_service/repositories/adspace"

var (
	AdspaceRepo adspace.AdspaceRepoInterface = adspace.AdspaceRepoImpl{}
)
