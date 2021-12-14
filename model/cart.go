/*
 * @Author: devilswj
 * @Date: 2021-11-26 16:48:23
 * @LastEditors: devilswj
 * @LastEditTime: 2021-11-26 16:49:55
 * @FilePath: \shop\model\cart.go
 */
package model

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserID    uint
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `grom:"not null"`
	BossID    uint
	Num       uint
	MaxNum    uint
	Check     bool
}
