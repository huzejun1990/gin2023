// @Author huzejun 2023/12/26 6:18:00
package main

import (
	"fmt"
	"gin2023/routers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//fmt.Println(Sum(1,2))
	fmt.Println(LogMiddleware(LogMiddleware(Sum))(1, 2))
	//return
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	routers.InitRouters(r)

	v2 := r.Group("/v2")
	v1 := r.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v2.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Sum(a, b int) int {
	return a + b
}

func LogMiddleware(in func(a, b int) int) func(a, b int) int {
	return func(a, b int) int {
		log.Println("called LogMiddleware")
		return in(a, b)
	}
}
