/*
 * @Author: devilswj
 * @Date: 2021-11-26 16:50:28
 * @LastEditors: devilswj
 * @LastEditTime: 2021-11-26 16:57:28
 * @FilePath: \shop\model\Order.go
 */
package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserID"`
	UserID    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeignKey:BossID"`
	BossID    uint    `gorm:"not null"`
	Address   Address `gorm:"ForeignKey:AddressID"`
	AddressID uint    `gorm:"notnull"`
	Num       uint
	OrderNum  int64
	Type      uint
	Money     int
}
