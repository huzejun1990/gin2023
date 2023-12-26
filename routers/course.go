// @Author huzejun 2023/12/27 3:31:00
package routers

import (
	"gin2023/course"
	"github.com/gin-gonic/gin"
)

func initCourse(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		//路径值参
		v1.GET("/course/:id", course.Get)
		//新增
		v1.POST("/course", course.Add)
		//编辑
		v1.PUT("/course", course.Update)
		//删除
		v1.DELETE("/course", course.Delete)

	}
}
