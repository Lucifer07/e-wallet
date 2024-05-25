package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/handler"
	middleware "github.com/Lucifer07/e-wallet/middleware"
	"github.com/Lucifer07/e-wallet/mocks"
	"github.com/Lucifer07/e-wallet/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPasswordTokenHandler(t *testing.T) {
	t.Run("should return password token handler if succes", func(t *testing.T) {
		// given
		targetHandler := handler.PasswordTokenHandler{}
		mockService := new(mocks.PasswordTokenService)
		// when
		handler := handler.NewPasswordTokenHandler(mockService)
		assert.IsType(t, targetHandler, *handler)
	})
}

func TestPasswordTokenHandler_CreateToken(t *testing.T) {
	t.Run("shoul return 200 if success", func(t *testing.T) {
		// given
		r := gin.New()
		r.Use(middleware.CustomMiddlewareError)
		mockPassword := new(mocks.PasswordTokenService)
		handlerPassword := handler.NewPasswordTokenHandler(mockPassword)
		reqBody, _ := json.Marshal(dto.GettokenRequest{
			Email: "test@tes.com",
		})
		// when
		mockPassword.On("CreateResetPassword", mock.Anything, mock.Anything).Return(nil, nil)
		r.POST("/get-token", handlerPassword.CreateToken)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/get-token", bytes.NewBuffer(reqBody))
		r.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("shoul return 400 if bad request", func(t *testing.T) {
		// given
		r := gin.New()
		r.Use(middleware.CustomMiddlewareError)
		mockPassword := new(mocks.PasswordTokenService)
		handlerPassword := handler.NewPasswordTokenHandler(mockPassword)
		reqBody, _ := json.Marshal([]dto.GettokenRequest{
			dto.GettokenRequest{
				Email: "test@tes.com",
			},
		})
		// when
		mockPassword.On("CreateResetPassword", mock.Anything, mock.Anything).Return(nil, nil)
		r.POST("/get-token", handlerPassword.CreateToken)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/get-token", bytes.NewBuffer(reqBody))
		r.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("shoul return 500 if internal server error", func(t *testing.T) {
		// given
		r := gin.New()
		r.Use(middleware.CustomMiddlewareError)
		mockPassword := new(mocks.PasswordTokenService)
		handlerPassword := handler.NewPasswordTokenHandler(mockPassword)
		reqBody, _ := json.Marshal(dto.GettokenRequest{
			Email: "test@tes.com",
		})
		// when
		mockPassword.On("CreateResetPassword", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		r.POST("/get-token", handlerPassword.CreateToken)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/get-token", bytes.NewBuffer(reqBody))
		r.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestPasswordTokenHandler_ResetPassword(t *testing.T) {
	t.Run("shoul return 204 if success", func(t *testing.T) {
		// given
		r := gin.New()
		r.Use(middleware.CustomMiddlewareError)
		mockPassword := new(mocks.PasswordTokenService)
		handlerPassword := handler.NewPasswordTokenHandler(mockPassword)
		reqBody, _ := json.Marshal(dto.TokenPassword{
			Token:    "ufehHgNrwbHpEKD0",
			Password: "test",
		})
		// when
		mockPassword.On("ResetPassword", mock.Anything, mock.Anything).Return( nil)
		r.POST("/reset-password", handlerPassword.ResetPassword)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(reqBody))
		r.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
	t.Run("shoul return 400 if bad request", func(t *testing.T) {
		// given
		r := gin.New()
		r.Use(middleware.CustomMiddlewareError)
		mockPassword := new(mocks.PasswordTokenService)
		handlerPassword := handler.NewPasswordTokenHandler(mockPassword)
		reqBody, _ := json.Marshal(dto.TokenPassword{
			Token:    "ufehHgNrwbHpE",
			Password: "test",
		})
		// when
		mockPassword.On("ResetPassword", mock.Anything, mock.Anything).Return(nil)
		r.POST("/reset-password", handlerPassword.ResetPassword)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(reqBody))
		r.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("should return 500 if internal server error", func(t *testing.T) {
		// given
		r := gin.New()
		r.Use(middleware.CustomMiddlewareError)
		mockPassword := new(mocks.PasswordTokenService)
		handlerPassword := handler.NewPasswordTokenHandler(mockPassword)
		reqBody, _ := json.Marshal(dto.TokenPassword{
			Token:    "ufehHgNrwbHpEKD0",
			Password: "test",
		})
		// when
		mockPassword.On("ResetPassword", mock.Anything, mock.Anything).Return(util.ErrorInternal)
		r.POST("/reset-password", handlerPassword.ResetPassword)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(reqBody))
		r.ServeHTTP(w, req)
		// then
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
