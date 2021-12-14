/*
 * @Author: devilswj
 * @Date: 2021-11-26 17:01:32
 * @LastEditors: devilswj
 * @LastEditTime: 2021-11-29 08:57:51
 * @FilePath: \shop\model\category.go
 */

package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CategoryName string
}
