package auctions

import (
	"PERSONAL/ad_space_auction_service/controllers"
	"PERSONAL/ad_space_auction_service/providers"
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
