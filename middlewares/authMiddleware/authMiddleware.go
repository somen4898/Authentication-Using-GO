package authMiddleware

import (
	"github.com/gin-gonic/gin"
	"go-auth/core/results"
	"go-auth/helpers/jwtHelper"
	"net/http"
	"strings"
)

func GetAccessToRoute(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		_token := c.Request.Header["Authorization"]
		if _token != nil && strings.Contains(_token[0], "Bearer") {
			token, err := jwtHelper.Parse(_token[0][len("Bearer "):])
			if err != nil {
				c.JSON(http.StatusInternalServerError, results.Error(err.Error()))
				c.Abort()
				return
			}
			if token.Valid {
				endpoint(c)
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, results.Error("Not Authorized"))
			c.Abort()
			return
		}
	}
}
