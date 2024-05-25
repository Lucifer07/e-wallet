package middleware

import (
	"net/http"
	"strconv"

	"github.com/Lucifer07/e-wallet/response"
	"github.com/Lucifer07/e-wallet/util"
	"github.com/gin-gonic/gin"
)

func MiddlewareJWTAuthorization(c *gin.Context) {
	const bearer = "Bearer "
	token := ""
	header := c.GetHeader("Authorization")
	if header != "" {
		token = header[len(bearer):]
	}
	if token != "" {
		claims, err := util.ParseAndVerify(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseMsgErr{Message: util.ErrorInvalidToken.Error()})
			return
		}
		if claims != nil {
			data := make(map[string]string, 0)
			id := int(claims["id"].(float64))
			if id >= 1 {
				data["id"] = strconv.Itoa(id)

				data["email"] = claims["email"].(string)
				
				c.Set("data", data)
				c.Next()
				return
			}
		}
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseMsgErr{Message: util.ErrorUnauthorized.Error()})
}
