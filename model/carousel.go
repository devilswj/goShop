/*
 * @Author: devilswj
 * @Date: 2021-11-29 09:00:38
 * @LastEditors: devilswj
 * @LastEditTime: 2021-11-29 09:02:46
 * @FilePath: \shop\model\carousel.go
 */
package model

import "github.com/jinzhu/gorm"

type Carousel struct {
	gorm.Model
	ImgPath   string
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `gorm:"not null"`
}
