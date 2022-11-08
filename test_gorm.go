//go:build linux
// +build linux

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model //代表继承关系
	Code       string

	Price uint
}

func create(db *gorm.DB) {
	//创建表
	db.AutoMigrate(&Product{})
}
func insert(db *gorm.DB) {

	p := Product{
		Code:  "1001",
		Price: 100,
	}
	//插入数据
	db.Create(&p)
}

// 查找
func find(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	fmt.Printf("p:%v\n", p)
	db.First(&p, "code=?", "1001") //查找code=1001 的数据
	fmt.Printf("p:%v\n", p)
}

// 更新
func update(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	db.Model(&p).Update("price", 1000)
	db.Model(&p).Updates(Product{Price: 1003, Code: "1002"})
	db.Model(&p).Updates(map[string]interface{}{"Price": 1006, "Code": 1005})
}
func del(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	db.Delete(&p, 1) /// 并不是物理删除  只是一个软删除  添加一个删除标记
}
func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	//find(db)
	//update(db)
	del(db)
}
