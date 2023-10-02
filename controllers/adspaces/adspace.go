package adspaces

import (
	"net/http"

	"github.com/samsunil1999/ad_space_auction_service/controllers"
	"github.com/samsunil1999/ad_space_auction_service/providers"
	"github.com/samsunil1999/ad_space_auction_service/transformers"
	"github.com/samsunil1999/ad_space_auction_service/validators"

	"github.com/gin-gonic/gin"
)

func CreateAdspace(ctx *gin.Context) {
	req, err := validators.ValiateCreateAdspaceReq(ctx)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusUnprocessableEntity, &controllers.ErrorResp{
			Code:    1,
			Message: "Invali request, " + err.Error(),
		})
		return
	}

	resp, err := providers.AdspaceSvc.CreateAdspace(req)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, transformers.GetAdspaceModel(resp))
}

func ListAllAdspaces(ctx *gin.Context) {
	resp, err := providers.AdspaceSvc.GetAllAvailableAdspace()
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, resp)

}

func GetAdspaceById(ctx *gin.Context) {
	id, err := validators.ValidateGetAdspaceById(ctx)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusUnprocessableEntity, &controllers.ErrorResp{
			Code:    1,
			Message: "Invali request, " + err.Error(),
		})
		return
	}

	resp, err := providers.AdspaceSvc.GetAdspaceById(id)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, transformers.GetAdspaceModel(resp))
}

func UpdateAdspaceById(ctx *gin.Context) {
	req, id, err := validators.ValidateUpdateAdspaceByIdReq(ctx)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusUnprocessableEntity, &controllers.ErrorResp{
			Code:    1,
			Message: "Invali request, " + err.Error(),
		})
		return
	}

	resp, err := providers.AdspaceSvc.UpdateAdspaceById(id, req)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, transformers.GetAdspaceModel(resp))
}

func DeleteAdspace(ctx *gin.Context) {
	id, err := validators.ValidateGetAdspaceById(ctx)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusUnprocessableEntity, &controllers.ErrorResp{
			Code:    1,
			Message: "Invali request, " + err.Error(),
		})
		return
	}

	resp, err := providers.AdspaceSvc.DeleteAdspaceById(id)
	if err != nil {
		controllers.ReturnJsonStruct(ctx, http.StatusInternalServerError, &controllers.ErrorResp{
			Code:    2,
			Message: "Something went wrong, " + err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, http.StatusOK, resp)
}
