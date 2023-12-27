// @Author huzejun 2023/12/27 21:25:00
package middleware

import (
	"gin2023/jwt_plugin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//权限校验
		accessToken := c.Request.Header.Get("access_token")
		data := &jwt_plugin.Data{}
		err := jwt_plugin.Verify(accessToken, data)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "身份认证失败",
			})
			c.Abort()
		}
		c.Set("auth_info", data)
		c.Next()
		//fmt.Println(accessToken)
	}
}
