/*
 * @Author: devilswj
 * @Date: 2021-12-22 19:35:02
 * @LastEditors: devilswj
 * @LastEditTime: 2021-12-22 19:40:42
 * @FilePath: \shop\api\user.go
 */
package api

import (
	"shop/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserRegisterService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
