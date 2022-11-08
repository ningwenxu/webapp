package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db = d
}
func test1() {
	db.Session(&gorm.Session{DryRun: true})
}

// one-many
func main() {
	//type Language struct {
	//	gorm.Model
	//	Name string
	//}
	//type User struct {
	//	gorm.Model
	//	Languages []Language `gorm:"many2many:user_languages"`
	//}
	//
	//db.AutoMigrate(&User{}, &Language{}) //这两个有先后顺序
	//l := Language{
	//	Name: "a",
	//}
	//l2 := Language{
	//	Name: "b",
	//}
	//user := User{
	//	Languages: []Language{l, l2},
	//}
	////db.Create(&l)
	////db.Create(&l2)
	//db.Create(&user)
}
