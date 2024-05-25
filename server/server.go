package server

import (
	middleware "github.com/Lucifer07/e-wallet/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler RouterOpt) *gin.Engine {
	r := gin.New()
	r.MaxMultipartMemory = 8 << 20
	r.Static("/public/images/", "images/")
	r.ContextWithFallback = true
	r.Use(
		middleware.Log,
		gin.Recovery(),
		middleware.CustomMiddlewareError,
	)
	return Route(r, handler)
}
