/*
 * @Author: devilswj
 * @Date: 2021-11-26 16:35:15
 * @LastEditors: devilswj
 * @LastEditTime: 2021-11-26 16:41:43
 * @FilePath: \shop\model\ProductImg.go
 */
package model

import "github.com/jinzhu/gorm"

// 商品图片表

type ProductImg struct {
	gorm.Model
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `grom:"not null"`
	ImgPath   string
}
