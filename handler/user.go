package handler

import (
	"log"
	"net/http"

	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/response"
	"github.com/Lucifer07/e-wallet/service"
	"github.com/Lucifer07/e-wallet/util"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewuserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
func (h *UserHandler) Login(ctx *gin.Context) {
	var user dto.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	jwt, err := h.userService.Login(ctx, user)
	if err != nil {
		log.Println(err)
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success, Data: jwt})
}
func (h *UserHandler) Register(ctx *gin.Context) {
	var user dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	registerData, err := h.userService.Register(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success, Data: registerData})
}
func (h *UserHandler) GetSelf(ctx *gin.Context) {
	data, err := h.userService.GetSelf(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success, Data: data})
}
func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	var user dto.UpdateProfile
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	err := h.userService.UpdateProfile(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusNoContent, "")
}
func (h *UserHandler) UpdateAvatar(ctx *gin.Context) {
	err := h.userService.UpdateAvatar(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusNoContent, "")
}
