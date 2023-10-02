package auctions

import (
	"PERSONAL/ad_space_auction_service/controllers"
	"PERSONAL/ad_space_auction_service/providers"
	"PERSONAL/ad_space_auction_service/transformers"
	"PERSONAL/ad_space_auction_service/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllLiveAuctions(ctx *gin.Context) {
	resp, err := providers.AuctionSvc.GetAllLiveAuctions()
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, resp)
}

func GetAuctionByAdspaceId(ctx *gin.Context) {
	id, err := validators.ValidateGetAdspaceById(ctx)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusUnprocessableEntity, &controllers.ErrorResp{
			Code:    1,
			Message: "Invali request, " + err.Error(),
		})
		return
	}

	resp, err := providers.AuctionSvc.GetAuctionById(id)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, resp)
}

func BidOnAuction(ctx *gin.Context) {
	req, err := validators.ValidateBidOnAuctionReq(ctx)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusUnprocessableEntity, &controllers.ErrorResp{
			Code:    1,
			Message: "Invali request, " + err.Error(),
		})
		return
	}

	resp, err := providers.AuctionSvc.NewBidOnAuction(req)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, transformers.GetBidsModel(resp))
}
