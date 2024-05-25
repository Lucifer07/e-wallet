package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
	"github.com/Lucifer07/e-wallet/handler"
	middleware "github.com/Lucifer07/e-wallet/middleware"
	"github.com/Lucifer07/e-wallet/mocks"
	"github.com/Lucifer07/e-wallet/util"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewHistoryHandler(t *testing.T) {
	t.Run("should return handler if succes", func(t *testing.T) {
		// given
		targetHandler := handler.HistoryHandler{}
		mockService := new(mocks.HistoryService)
		// when
		handler := handler.NewHistoryHandler(mockService)
		assert.IsType(t, targetHandler, *handler)
	})
}

func TestHistoryHandler_MyTransactions(t *testing.T) {
	t.Run("should return 200 if success", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		TargetData := []entity.HistoryTransaction{entity.HistoryTransaction{}}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		// when
		mockService.On("MyTransactions", mock.Anything, mock.Anything).Return(&TargetData, nil)
		route.GET("/user/transaction", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.MyTransactions(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/user/transaction", nil)
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return 500 if internal server error", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		// when
		mockService.On("MyTransactions", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		route.GET("/user/transaction", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.MyTransactions(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/user/transaction", nil)
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
func TestHistoryHandler_TopUpBank(t *testing.T) {
	t.Run("should return 200 if succes", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(dto.TopupBankRequest{AccountNumber: 1231, Amount: decimal.NewFromInt(100), Description: "test"})
		// when
		mockService.On("TopupBank", mock.Anything, mock.Anything).Return(nil, nil)
		route.POST("/user/topup/bank", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.TopUpBank(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/topup/bank", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return 400 if have bad request", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(
			[]dto.TopupBankRequest{
				dto.TopupBankRequest{AccountNumber: 1231, Amount: decimal.NewFromInt(100), Description: "test"},
			})
		// when
		route.POST("/user/topup/bank", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.TopUpBank(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/topup/bank", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("should return 500 if internal server error", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(dto.TopupBankRequest{AccountNumber: 1231, Amount: decimal.NewFromInt(100), Description: "test"})
		// when
		mockService.On("TopupBank", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		route.POST("/user/topup/bank", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.TopUpBank(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/topup/bank", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHistoryHandler_TopUpCC(t *testing.T) {
	t.Run("should return 200 if succes", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(dto.TopupCreditCardRequest{CCNumber: 1231, Amount: decimal.NewFromInt(100), Description: "test"})
		// when
		mockService.On("TopupCreditCard", mock.Anything, mock.Anything).Return(nil, nil)
		route.POST("/user/topup/creditcard", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.TopUpCC(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/topup/creditcard", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return 400 if have bad request", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(
			[]dto.TopupCreditCardRequest{
				dto.TopupCreditCardRequest{CCNumber: 1231, Amount: decimal.NewFromInt(100), Description: "test"},
			})

		// when
		route.POST("/user/topup/creditcard", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.TopUpCC(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/topup/creditcard", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("should return 500 if internal server error", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(dto.TopupCreditCardRequest{CCNumber: 1231, Amount: decimal.NewFromInt(100), Description: "test"})
		// when
		mockService.On("TopupCreditCard", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		route.POST("/user/topup/creditcard", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.TopUpCC(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/topup/creditcard", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHistoryHandler_TopUpPaylater(t *testing.T) {
	t.Run("should return 200 if succes", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(dto.TopupPaylaterRequest{Amount: decimal.NewFromInt(100), Description: "test"})
		// when
		mockService.On("TopupPayLater", mock.Anything, mock.Anything).Return(nil, nil)
		route.POST("/user/topup/paylater", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.TopUpPaylater(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/topup/paylater", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return 400 if have bad request", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(
			[]dto.TopupPaylaterRequest{
				dto.TopupPaylaterRequest{Amount: decimal.NewFromInt(100), Description: "test"},
			})
		// when
		route.POST("/user/topup/paylater", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.TopUpPaylater(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/topup/paylater", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("should return 500 if internal server error", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(dto.TopupPaylaterRequest{Amount: decimal.NewFromInt(100), Description: "test"})
		// when
		mockService.On("TopupPayLater", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		route.POST("/user/topup/paylater", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.TopUpPaylater(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/topup/paylater", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHistoryHandler_Transfer(t *testing.T) {
	t.Run("should return 200 if succes", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(dto.TransferRequest{WalletNumber: "xxx" ,Amount: decimal.NewFromInt(100), Description: "test"})
		// when
		mockService.On("Transfer", mock.Anything, mock.Anything).Return(nil, nil)
		route.POST("/user/transfer", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.Transfer(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/transfer", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return 400 if have bad request", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(
			[]dto.TransferRequest{
				dto.TransferRequest{WalletNumber: "xxx" ,Amount: decimal.NewFromInt(100), Description: "test"},
			})
		// when
		route.POST("/user/transfer", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.Transfer(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/transfer", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("should return 500 if internal server error", func(t *testing.T) {
		// given
		data := map[string]string{"id": "1", "email": "test@test.com"}
		route := gin.New()
		route.Use(middleware.CustomMiddlewareError)
		mockService := new(mocks.HistoryService)
		handlerApp := handler.NewHistoryHandler(mockService)
		reqBody, _ := json.Marshal(dto.TransferRequest{WalletNumber: "xxx" ,Amount: decimal.NewFromInt(100), Description: "test"})
		// when
		mockService.On("Transfer", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		route.POST("/user/transfer", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.Transfer(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/transfer", bytes.NewBuffer(reqBody))
		route.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
