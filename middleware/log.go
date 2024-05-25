package middleware

import (
	"strconv"

	"github.com/Lucifer07/e-wallet/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Log(c *gin.Context) {
	log := util.Log()
	log.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	}).Info("Request received")
	c.Next()
	statusOk := 2
	statusFirst := strconv.Itoa(c.Writer.Status())
	statusCode, _ := strconv.Atoi(string(statusFirst[0]))
	if statusCode == statusOk {
		log.WithFields(logrus.Fields{
			"status": c.Writer.Status(),
		}).Info("Request completed")
		return
	}
	if len(c.Errors) > 0 {
		log.WithFields(logrus.Fields{
			"error":  c.Errors[0].Err,
			"status": c.Writer.Status(),
		}).Error("Request completed")
		return
	}
	log.WithFields(logrus.Fields{
		"status": c.Writer.Status(),
	}).Error("Request completed")
}
