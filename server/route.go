package server

import (
	"net/http"
	"net/http/pprof"

	middleware "github.com/Lucifer07/e-wallet/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Route(route *gin.Engine, handler RouterOpt) *gin.Engine {
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	route.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
	route.POST("/login", handler.UserHandler.Login)
	route.POST("/register", handler.UserHandler.Register)
	routeReset := route.Group("/reset")
	routeReset.POST("/get-token", handler.passwordHandler.CreateToken)
	routeReset.POST("/reset-password", handler.passwordHandler.ResetPassword)
	routeUser := route.Group("/user", middleware.MiddlewareJWTAuthorization)
	routeUser.GET("", handler.UserHandler.GetSelf)
	routeUser.POST("", handler.UserHandler.UpdateProfile)
	routeUser.POST("/avatar", handler.UserHandler.UpdateAvatar)
	routeUser.GET("/transaction", handler.historyHandler.MyTransactions)
	routeUser.POST("/transfer", handler.historyHandler.Transfer)
	routerTopup := routeUser.Group("/topup")
	routerTopup.POST("/bank", handler.historyHandler.TopUpBank)
	routerTopup.POST("/creditcard", handler.historyHandler.TopUpCC)
	routerTopup.POST("/paylater", handler.historyHandler.TopUpPaylater)
	route.GET("/metrics",gin.WrapH(promhttp.Handler()))
	route.GET("/debug/pprof/", gin.WrapH(http.HandlerFunc(pprof.Index)))
	route.GET("/debug/pprof/profile", gin.WrapH(http.HandlerFunc(pprof.Profile)))
	route.GET("/debug/pprof/heap", gin.WrapH(http.HandlerFunc(pprof.Handler("heap").ServeHTTP)))
	route.GET("/debug/pprof/block", gin.WrapH(http.HandlerFunc(pprof.Handler("block").ServeHTTP)))
	route.GET("/debug/pprof/goroutine", gin.WrapH(http.HandlerFunc(pprof.Handler("goroutine").ServeHTTP)))
	route.Run()
	return route
}
