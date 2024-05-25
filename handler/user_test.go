package handler_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/Lucifer07/e-wallet/dto"
// 	"github.com/Lucifer07/e-wallet/handler"
// 	middleware "github.com/Lucifer07/e-wallet/middleware"
// 	"github.com/Lucifer07/e-wallet/mocks"
// 	"github.com/Lucifer07/e-wallet/util"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestUserHandler_Login(t *testing.T) {
// 	t.Run("should return 200 if success", func(t *testing.T) {
// 		// given
// 		r := gin.New()
// 		r.Use(middleware.CustomMiddlewareError)
// 		mockUser := new(mocks.UserService)
// 		handlerUser := handler.NewuserHandler(mockUser)
// 		reqBody, _ := json.Marshal(dto.Login{
// 			Email:    "test@test.com",
// 			Password: "test",
// 		})
// 		// when
// 		mockUser.On("Login", mock.Anything, mock.Anything).Return(nil, nil)
// 		r.POST("/login", handlerUser.Login)
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
// 		r.ServeHTTP(w, req)
// 		// then
// 		assert.Equal(t, http.StatusOK, w.Code)
// 	})
// 	t.Run("should return 400 if have bad request", func(t *testing.T) {
// 		// given
// 		r := gin.New()
// 		r.Use(middleware.CustomMiddlewareError)
// 		mockUser := new(mocks.UserService)
// 		handlerUser := handler.NewuserHandler(mockUser)
// 		reqBody, _ := json.Marshal(dto.Login{
// 			Email: "test@test.com",
// 		})
// 		// when
// 		mockUser.On("Login", mock.Anything, mock.Anything).Return(nil, nil)
// 		r.POST("/login", handlerUser.Login)
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
// 		r.ServeHTTP(w, req)
// 		// then
// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 	})
// 	t.Run("should return 500 if internal server error", func(t *testing.T) {
// 		// given
// 		r := gin.New()
// 		r.Use(middleware.CustomMiddlewareError)
// 		mockUser := new(mocks.UserService)
// 		handlerUser := handler.NewuserHandler(mockUser)
// 		reqBody, _ := json.Marshal(dto.Login{
// 			Email:    "test@test.com",
// 			Password: "test",
// 		})
// 		// when
// 		mockUser.On("Login", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
// 		r.POST("/login", handlerUser.Login)
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
// 		r.ServeHTTP(w, req)
// 		// then
// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 	})
// }

// func TestUserHandler_Register(t *testing.T) {
// 	t.Run("should return 200 if success", func(t *testing.T) {
// 		// given
// 		r := gin.New()
// 		r.Use(middleware.CustomMiddlewareError)
// 		mockUser := new(mocks.UserService)
// 		handlerUser := handler.NewuserHandler(mockUser)
// 		reqBody, _ := json.Marshal(dto.RegisterRequest{
// 			Name:     "test",
// 			Email:    "test@test.com",
// 			Password: "test",
// 		})
// 		// when
// 		mockUser.On("Register", mock.Anything, mock.Anything).Return(nil, nil)
// 		r.POST("/register", handlerUser.Register)
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
// 		r.ServeHTTP(w, req)
// 		// then
// 		assert.Equal(t, http.StatusOK, w.Code)
// 	})
// 	t.Run("should return 400 if have a bad request", func(t *testing.T) {
// 		// given
// 		r := gin.New()
// 		r.Use(middleware.CustomMiddlewareError)
// 		mockUser := new(mocks.UserService)
// 		handlerUser := handler.NewuserHandler(mockUser)
// 		reqBody, _ := json.Marshal(dto.RegisterRequest{
// 			Name:  "test",
// 			Email: "test@test.com",
// 		})
// 		// when
// 		mockUser.On("Register", mock.Anything, mock.Anything).Return(nil, nil)
// 		r.POST("/register", handlerUser.Register)
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
// 		r.ServeHTTP(w, req)
// 		// then
// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 	})
// 	t.Run("should return 500 if internal server error", func(t *testing.T) {
// 		// given
// 		r := gin.New()
// 		r.Use(middleware.CustomMiddlewareError)
// 		mockUser := new(mocks.UserService)
// 		handlerUser := handler.NewuserHandler(mockUser)
// 		reqBody, _ := json.Marshal(dto.RegisterRequest{
// 			Name:     "test",
// 			Email:    "test@test.com",
// 			Password: "test",
// 		})
// 		// when
// 		mockUser.On("Register", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
// 		r.POST("/register", handlerUser.Register)
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
// 		r.ServeHTTP(w, req)
// 		// then
// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 	})
// }

// func TestUserHandler_GetSelf(t *testing.T) {
// 	t.Run("should return 200 if succes", func(t *testing.T) {
// 		// given
// 		r := gin.New()
// 		r.Use(middleware.CustomMiddlewareError)
// 		mockUser := new(mocks.UserService)
// 		handlerUser := handler.NewuserHandler(mockUser)
// 		// when
// 		mockUser.On("GetSelf", mock.Anything).Return(nil, nil)
// 		r.POST("/user", handlerUser.GetSelf)
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/user", nil)
// 		r.ServeHTTP(w, req)
// 		// then
// 		assert.Equal(t, http.StatusOK, w.Code)
// 	})
// 	t.Run("should return 500 if internal server error", func(t *testing.T) {
// 		// given
// 		r := gin.New()
// 		r.Use(middleware.CustomMiddlewareError)
// 		mockUser := new(mocks.UserService)
// 		handlerUser := handler.NewuserHandler(mockUser)
// 		// when
// 		mockUser.On("GetSelf", mock.Anything).Return(nil, util.ErrorInternal)
// 		r.POST("/user", handlerUser.GetSelf)
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/user", nil)
// 		r.ServeHTTP(w, req)
// 		// then
// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 	})
// }
