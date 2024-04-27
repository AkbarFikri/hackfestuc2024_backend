package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/pkg/helper"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/pkg/model"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/service"

)

type AqiHandler struct {
	AqiService service.AqiService
}

func NewAqi(as service.AqiService) AqiHandler {
	return AqiHandler{
		AqiService: as,
	}
}

func (h *AqiHandler) GetCurrentPosition(ctx *gin.Context) {
	var param model.AqiParam

	if err := ctx.ShouldBindQuery(&param); err != nil {
		helper.ErrorResponse(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid request payload",
		})
		return
	}

	data, err := h.AqiService.FetchAqiData(param)
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)
}
