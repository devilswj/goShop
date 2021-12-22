/*
 * @Author: devilswj
 * @Date: 2021-11-26 15:54:25
 * @LastEditors: devilswj
 * @LastEditTime: 2021-12-22 19:33:02
 * @FilePath: \shop\main.go
 */
package main

import (
	"shop/conf"
	"shop/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	routes.Init(r)
}
