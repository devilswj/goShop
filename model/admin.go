/*
 * @Author: devilswj
 * @Date: 2021-11-29 08:58:05
 * @LastEditors: devilswj
 * @LastEditTime: 2021-11-29 08:59:45
 * @FilePath: \shop\model\admin.go
 */
package model

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}
