package bidders

import (
	"PERSONAL/ad_space_auction_service/controllers"
	"PERSONAL/ad_space_auction_service/providers"
	"PERSONAL/ad_space_auction_service/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterBidder(ctx *gin.Context) {
	req, err := validators.ValidateRegisterBidderReq(ctx)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusUnprocessableEntity, &controllers.ErrorResp{
			Code:    1,
			Message: "Invali request, " + err.Error(),
		})
		return
	}

	resp, err := providers.BidderSvc.RegisterBidder(req)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, resp)
}
