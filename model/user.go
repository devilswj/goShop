/*
 * @Author: devilswj
 * @Date: 2021-11-26 16:22:31
 * @LastEditors: devilswj
 * @LastEditTime: 2021-11-26 16:41:09
 * @FilePath: \shop\model\user.go
 */
package model

import "github.com/jinzhu/gorm"

// 用户表

type User struct {
	gorm.Model
	UserName       string `grom:"unique"`
	Email          string
	PasswordDigest string
	Nickname       string `gorm:"not null"`
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          int
}
