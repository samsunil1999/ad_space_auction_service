package main

import (
	"PERSONAL/ad_space_auction_service/configs"
	"PERSONAL/ad_space_auction_service/controllers/adspaces"
	"PERSONAL/ad_space_auction_service/controllers/auctions"
	"PERSONAL/ad_space_auction_service/controllers/bidders"
	"PERSONAL/ad_space_auction_service/database"
	"log"

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
	router.GET("/bidders", bidders.RegisterBidder) // list all bidders
	router.GET("/bidders/{id}")                    // get bidder by id
	router.POST("/bidders")                        // register new bidder
	router.PUT("/bidders/{id}")                    // update bidder details
	router.DELETE("/bidders/{id}")                 // delete bidder

	// auction endpoints
	router.GET("/auctions/list-all", auctions.GetAllLiveAuctions)        // list all live auctions
	router.GET("/auctions/adspaces/:id", auctions.GetAuctionByAdspaceId) // Retrieve details of a specific auction, including the winning bid if the auction has ended & active bids if the auction is live
	router.POST("/auctions/adspaces/:id/bid", auctions.BidOnAuction)     // Allow bidders to submit bids for an ongoing auction.
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
