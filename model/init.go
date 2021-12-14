/*
 * @Author: devilswj
 * @Date: 2021-11-26 16:14:50
 * @LastEditors: devilswj
 * @LastEditTime: 2021-12-14 09:32:23
 * @FilePath: \shop\model\init.go
 */
package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true) // Gorm 打印
	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)      // 默认不加复数s
	db.DB().SetMaxIdleConns(20) //设置连接池
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}

func migration() {
	DB.Set("gorm:table_options", "chartset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Product{}).
		AutoMigrate(&Carousel{}).
		AutoMigrate(&Category{}).
		AutoMigrate(&Favorite{}).
		AutoMigrate(&ProductImg{}).
		AutoMigrate(&ProductInfoImg{}).
		AutoMigrate(&ProductParamImg{}).
		AutoMigrate(&Order{}).
		AutoMigrate(&Cart{}).
		AutoMigrate(&Admin{}).
		AutoMigrate(&Address{})
	DB.Model(&Cart{}).AddForeignKey("product_id", "Product(id)", "CASCADE", "CASCADE")
	DB.Model(&Order{}).AddForeignKey("user_id", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&Order{}).AddForeignKey("address_id", "Address(id)", "CASCADE", "CASCADE")
	DB.Model(&Order{}).AddForeignKey("product_id", "Product(id)", "CASCADE", "CASCADE")
	DB.Model(&Order{}).AddForeignKey("boss_id", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&Favorite{}).AddForeignKey("boss_id", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&Favorite{}).AddForeignKey("user_id", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&Favorite{}).AddForeignKey("product_id", "Product(id)", "CASCADE", "CASCADE")
	DB.Model(&Product{}).AddForeignKey("category_id", "Category(id)", "CASCADE", "CASCADE")
	DB.Model(&ProductImg{}).AddForeignKey("product_id", "Product(id)", "CASCADE", "CASCADE")
	DB.Model(&ProductInfoImg{}).AddForeignKey("product_id", "Product(id)", "CASCADE", "CASCADE")
	DB.Model(&ProductParamImg{}).AddForeignKey("product_id", "Product(id)", "CASCADE", "CASCADE")
	DB.Model(&Address{}).AddForeignKey("user_id", "User(id)", "CASCADE", "CASCADE")
}
