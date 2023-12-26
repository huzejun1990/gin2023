// @Author huzejun 2023/12/27 4:32:00
package course

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type course struct {
	Name     string
	Teacher  string
	Duration int
}

func Add(c *gin.Context) {
	req := &course{}
	err := c.ShouldBind(req)
	err := c.Bind(req)

	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	})
}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	})
}

func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	})
}
