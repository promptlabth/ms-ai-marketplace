package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
		
		path := c.Request.URL.Path
		println("path: ",path)
		segments := strings.Split(path, "/")
		println("segments: ",segments)
		if len(segments) > 1 {
			lang := segments[1]
			println("lang: ",lang)
			switch lang {
			case "en":
				c.Set("language", "en")
			default:
				c.Set("language", "th")
			}
		} else {
			c.Set("language", "th")
		}

		c.Next()

	}
}
