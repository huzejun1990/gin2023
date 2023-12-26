// @Author huzejun 2023/12/26 6:40:00
package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	api := r.Group("/api")
	//课程相关接口
	initCourse(api)
	//用户相关接口
	initUser(api)
	//登陆
	initLogin(api)
}
