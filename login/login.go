// @Author huzejun 2023/12/27 21:45:00
package login

import (
	"gin2023/jwt_plugin"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	data := jwt_plugin.Data{
		Name:   "nick",
		Age:    18,
		Gender: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	sign, err := jwt_plugin.Sign(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": sign,
	})
}
