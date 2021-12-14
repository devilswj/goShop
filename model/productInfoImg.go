/*
 * @Author: devilswj
 * @Date: 2021-11-26 16:46:08
 * @LastEditors: devilswj
 * @LastEditTime: 2021-12-14 09:31:50
 * @FilePath: \shop\model\productInfoImg.go
 */
package model

import "github.com/jinzhu/gorm"

type ProductInfoImg struct {
	gorm.Model
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `grom:"not null"`
	ImgPath   string
}
