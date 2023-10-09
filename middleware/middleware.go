package middleware

import (
	"net/http"

	"github.com/1412270/ecommerce-golang/tokens"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClienToken := c.Request.Header.Get("token")
		if ClienToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no authorization header"})
			c.Abort()
			return
		}

		claims, err := tokens.ValidateToken(ClienToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
