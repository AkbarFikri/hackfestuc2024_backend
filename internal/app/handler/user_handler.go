package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/Aironment-BE/internal/app/service"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/helper"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/model"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUser(us service.UserService) UserHandler {
	return UserHandler{
		UserService: us,
	}
}

func (h *UserHandler) CurrentUser(ctx *gin.Context) {
	user := helper.GetUserLoginData(ctx)

	data, err := h.UserService.Current(user)
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)
}

func (h *UserHandler) GetAirqualityPoints(ctx *gin.Context) {
	var param model.AqiParam

	if err := ctx.ShouldBindQuery(&param); err != nil {
		helper.ErrorResponse(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid request payload",
		})
		return
	}

	data, err := h.UserService.FetchAirQualitysPoints(param)
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)
}
