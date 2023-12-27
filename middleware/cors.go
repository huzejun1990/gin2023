// @Author huzejun 2023/12/27 20:25:00
package middleware

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors1() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("Called Cors1")
	}
}

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowHeaders: []string{
				"Origin", "Content-Length", "Content-Type",
			},
			AllowMethods: []string{
				"GET", "POST", "DELETE", "PUT", "HEAD", "OPTIONS",
			},
		})
}
