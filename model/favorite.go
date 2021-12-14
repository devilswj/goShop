/*
 * @Author: devilswj
 * @Date: 2021-12-14 08:57:46
 * @LastEditors: devilswj
 * @LastEditTime: 2021-12-14 09:03:46
 * @FilePath: \shop\model\favorite.go
 */
package model

import "github.com/jinzhu/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserID"`
	UserID    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeignKey:BossID"`
	BossID    uint    `gorm:"not null"`
}
