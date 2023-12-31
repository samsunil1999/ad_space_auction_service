package bidders

import (
	"net/http"

	"github.com/samsunil1999/ad_space_auction_service/controllers"
	"github.com/samsunil1999/ad_space_auction_service/providers"
	"github.com/samsunil1999/ad_space_auction_service/validators"

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

func GetBidderById(ctx *gin.Context) {
	id, err := validators.ValidateBidderId(ctx)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusUnprocessableEntity, &controllers.ErrorResp{
			Code:    1,
			Message: "Invali request, " + err.Error(),
		})
		return
	}

	resp, err := providers.BidderSvc.GetBidderById(id)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, resp)
}

func GetAllBidders(ctx *gin.Context) {
	resp, err := providers.BidderSvc.GetAllBidders()
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, resp)
}

func UpdateBidderDetails(ctx *gin.Context) {
	req, id, err := validators.ValidateUpdateBidderReq(ctx)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusUnprocessableEntity, &controllers.ErrorResp{
			Code:    1,
			Message: "Invali request, " + err.Error(),
		})
		return
	}

	resp, err := providers.BidderSvc.UpdateBidderDetails(id, req)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, resp)
}

func DeRegisterBidder(ctx *gin.Context) {
	id, err := validators.ValidateBidderId(ctx)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusUnprocessableEntity, &controllers.ErrorResp{
			Code:    1,
			Message: "Invali request, " + err.Error(),
		})
		return
	}

	resp, err := providers.BidderSvc.DeleteBidder(id)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, resp)
}
