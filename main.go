package main

import (
	"log"

	"github.com/samsunil1999/ad_space_auction_service/configs"
	"github.com/samsunil1999/ad_space_auction_service/controllers/adspaces"
	"github.com/samsunil1999/ad_space_auction_service/controllers/auctions"
	"github.com/samsunil1999/ad_space_auction_service/controllers/bidders"
	"github.com/samsunil1999/ad_space_auction_service/database"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	configs.LoadConfig()
	database.Connect()
	database.SyncDatabase()
}

func mapUrl() {
	// adspace endpoints
	router.GET("/adspaces/list-all", adspaces.ListAllAdspaces) // list all availale adspace
	router.GET("/adspaces/:id", adspaces.GetAdspaceById)       // get adspace by id
	router.POST("/adspaces", adspaces.CreateAdspace)           // create adspace
	router.PUT("/adspaces/:id", adspaces.UpdateAdspaceById)    // update adspace details
	router.DELETE("/adspaces/:id", adspaces.DeleteAdspace)     // delete adspace

	// bidder endpoints
	router.GET("/bidders/list-all", bidders.GetAllBidders)  // list all bidders
	router.GET("/bidders/:id", bidders.GetBidderById)       // get bidder by id
	router.POST("/bidders", bidders.RegisterBidder)         // register new bidder
	router.PUT("/bidders/:id", bidders.UpdateBidderDetails) // update bidder details
	router.DELETE("/bidders/:id", bidders.DeRegisterBidder) // delete bidder

	// auction endpoints
	router.GET("/auctions/list-all", auctions.GetAllLiveAuctions)        // list all live auctions
	router.GET("/auctions/adspaces/:id", auctions.GetAuctionByAdspaceId) // Retrieve details of a specific auction, including the winning bid if the auction has ended & active bids if the auction is live
	router.POST("/auctions/bid", auctions.BidOnAuction)                  // Allow bidders to submit bids for an ongoing auction.
}

func main() {
	// endpoint mapping
	mapUrl()

	// run the application
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}
