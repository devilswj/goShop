/*
 * @Author: devilswj
 * @Date: 2021-11-26 16:57:43
 * @LastEditors: devilswj
 * @LastEditTime: 2021-11-26 17:01:16
 * @FilePath: \shop\model\address.go
 */
package model

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserID"`
	UserID    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeignKey:BossID"`
	BossID    uint    `gorm:"not null"`
}
