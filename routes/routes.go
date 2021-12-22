/*
 * @Author: devilswj
 * @Date: 2021-12-22 19:26:49
 * @LastEditors: devilswj
 * @LastEditTime: 2021-12-22 19:30:21
 * @FilePath: \shop\routes\routes.go
 */
package routes

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register")
	}
}
