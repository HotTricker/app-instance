package middleware

import (
	"app-instance/internal/pkg/gojwt"
	"app-instance/internal/pkg/render"
	"strings"

	"github.com/gin-gonic/gin"
)

func ApiPriv() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		path := strings.TrimSpace(c.Request.URL.Path)
		if strings.Contains(path, "login") {
			c.Next()
			return
		}

		if len(auth) == 0 {
			respondWithError(c, render.CODE_ERR_NO_LOGIN, "no login")
			return
		}

		_, err := gojwt.ParseToken(auth)
		if err != nil {
			respondWithError(c, render.CODE_ERR_NO_LOGIN, "no login")
			return
		} else {
			c.Next()
		}

	}
}

func respondWithError(c *gin.Context, code int, message string) {
	render.CustomerError(c, code, message)
	c.Abort()
}
