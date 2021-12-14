package model

import "github.com/jinzhu/gorm"

// 商品参数表

type ProductParamImg struct {
	gorm.Model
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `grom:"not null"`
	ImgPath   string
}
