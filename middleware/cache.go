package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func CacheControl(cacheTTL int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "s-maxage="+strconv.Itoa(cacheTTL))

		c.Next()
	}
}
