package handler

import (
	"net/http"

	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/response"
	"github.com/Lucifer07/e-wallet/service"
	"github.com/Lucifer07/e-wallet/util"
	"github.com/gin-gonic/gin"
)

type PasswordTokenHandler struct {
	PasswordTokenService service.PasswordTokenService
}

func NewPasswordTokenHandler(PasswordTokenService service.PasswordTokenService) *PasswordTokenHandler {
	return &PasswordTokenHandler{
		PasswordTokenService: PasswordTokenService,
	}
}
func (h *PasswordTokenHandler) CreateToken(ctx *gin.Context) {
	var email dto.GettokenRequest
	if err := ctx.ShouldBindJSON(&email); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	responseData, err := h.PasswordTokenService.CreateResetPassword(ctx, email)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success, Data: responseData})
}
func (h *PasswordTokenHandler) ResetPassword(ctx *gin.Context) {
	var tokenPassword dto.TokenPassword
	if err := ctx.ShouldBindJSON(&tokenPassword); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	err := h.PasswordTokenService.ResetPassword(ctx, tokenPassword)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusNoContent, util.NoContent)
}
