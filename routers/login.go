// @Author huzejun 2023/12/27 3:45:00
package routers

import (
	"gin2023/login"
	"github.com/gin-gonic/gin"
)

func initLogin(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/login", login.Login)
	}
}
