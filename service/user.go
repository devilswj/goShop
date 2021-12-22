/*
 * @Author: devilswj
 * @Date: 2021-12-22 19:30:48
 * @LastEditors: devilswj
 * @LastEditTime: 2021-12-22 19:40:19
 * @FilePath: \shop\service\user.go
 */
package service

type UserRegisterService struct {
	Nickname string `form:"nickname" json:"nickname" bingding:"required,min=2,max=10"`
	UserName string `form:"username" json:"username" bingding:"required,min=5,max=15"`
	Password string `form:"password" json:"password" bingding:"required,min=8,max=16"`
}

func (service *UserRegisterService) Register() {

}
