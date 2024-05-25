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

type HistoryHandler struct {
	HistoryService service.HistoryService
}

func NewHistoryHandler(HistoryService service.HistoryService) *HistoryHandler {
	return &HistoryHandler{
		HistoryService: HistoryService,
	}
}
func (h *HistoryHandler) MyTransactions(ctx *gin.Context) {
	params := make(map[string]string, 0)
	params["page"] = ctx.Query("page")
	params["limit"] = ctx.Query("limit")
	params["description"] = ctx.Query("description")
	params["sortBy"] = ctx.Query("sortBy")
	params["sortOrder"] = ctx.Query("sortOrder")
	params["from"] = ctx.Query("from")
	params["to"] = ctx.Query("to")
	params["type"] = ctx.Query("type")
	data, err := h.HistoryService.MyTransactions(ctx, params)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success, Data: data})
}
func (h *HistoryHandler) TopUpBank(ctx *gin.Context) {
	var data dto.TopupBankRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {

		ctx.Error(util.ErrorBadRequest)
		return
	}
	responseData, err := h.HistoryService.TopupBank(ctx, data)
	if err != nil {
		log.Println("errr",err)
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success, Data: responseData})
}
func (h *HistoryHandler) TopUpCC(ctx *gin.Context) {
	var data dto.TopupCreditCardRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	responseData, err := h.HistoryService.TopupCreditCard(ctx, data)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success, Data: responseData})
}
func (h *HistoryHandler) TopUpPaylater(ctx *gin.Context) {
	var data dto.TopupPaylaterRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	responseData, err := h.HistoryService.TopupPayLater(ctx, data)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success, Data: responseData})
}
func (h *HistoryHandler) Transfer(ctx *gin.Context) {
	var data dto.TransferRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	responseData, err := h.HistoryService.Transfer(ctx, data)
	if err != nil {
		log.Println(err)
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success, Data: responseData})
}
