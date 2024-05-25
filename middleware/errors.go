package middleware

import (
	"net/http"

	"github.com/Lucifer07/e-wallet/response"
	"github.com/Lucifer07/e-wallet/util"
	"github.com/gin-gonic/gin"
)

func CustomMiddlewareError(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		firstError := c.Errors[0].Err
		errResponse := checkError(firstError)
		if errResponse.StatusCode != http.StatusInternalServerError {
			c.JSON(errResponse.StatusCode, errResponse)
			return
		}
		c.AbortWithStatusJSON(errResponse.StatusCode, errResponse)
		return
	}

}

func checkError(err error) response.ResponseMsgErr {
	switch err {
	case util.ErrorUserNotFound:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErrorUserNotFound.Error()}
	case util.ErrorWrongPassword:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErrorWrongPassword.Error()}
	case util.ErrorBadRequest:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErrorBadRequest.Error()}
	case util.ErrorEmailUnique:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErrorEmailUnique.Error()}
	case util.ErrorInvalidToken:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErrorInvalidToken.Error()}
	case util.ErrorBalance:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErrorBalance.Error()}
	case util.ErroMinimumTopUp:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErroMinimumTopUp.Error()}
	case util.ErroMinimumTranfer:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErroMinimumTranfer.Error()}
	case util.ErroMaximalTopUp:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErroMaximalTopUp.Error()}
	case util.ErroMaximalTranfer:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErroMaximalTranfer.Error()}
	case util.ErrorInvalidTransfer:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErrorInvalidTransfer.Error()}
	default:
		return response.ResponseMsgErr{StatusCode: http.StatusInternalServerError}
	}
}
