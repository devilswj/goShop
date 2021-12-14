/*
 * @Author: devilswj
 * @Date: 2021-11-26 16:26:57
 * @LastEditors: devilswj
 * @LastEditTime: 2021-11-26 16:41:48
 * @FilePath: \shop\model\product.go
 */

package model

import "github.com/jinzhu/gorm"

// 商品表

type Product struct {
	gorm.Model
	Name          string   `gorm:"size:255;index"`
	Category      Category `grom:"ForeignKey:CategoryID"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        int
	BossName      string
	BossAvatar    string
}
