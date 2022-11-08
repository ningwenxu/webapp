package main

//
//import (
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//var db *gorm.DB
//
//func init() {
//	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
//
//	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	if err != nil {
//		panic("failed to connect database")
//	}
//	db = d
//}
//
////type Company struct {
////	ID   int
////	Name string
////}
////type User struct {
////	gorm.Model
////	Name      string
////	CompanyID int
////	Company   Company
////}
//
//func main() {
//	type CreditCard struct {
//		gorm.Model
//		Number string
//		UserID uint
//	}
//	type User struct {
//		gorm.Model
//		CreditCard CreditCard
//	}
//	db.AutoMigrate(&User{}, &CreditCard{}) //这两个有先后顺序
//}
