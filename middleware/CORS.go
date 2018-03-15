package middleware

import "github.com/gin-gonic/gin"

func CORSMiddleware(whiteList []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		allowOrigin := ""
		for _, str := range whiteList {
			if str == c.Request.Header.Get("Origin") {
				allowOrigin = str
				break
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Accept-Language, X-CSRF-Token, Authorization, X-Requested-With, X-Access-Token, X-Lozi-Client, X-City-ID")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		} else {
			c.Next()
		}
	}
}
