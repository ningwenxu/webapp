package main

//
//import (
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"time"
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
//type User struct {
//	gorm.Model
//	Name     string
//	Age      int
//	Birthday time.Time //要大写
//	Active   bool
//}
//
//func createTable() {
//	db.AutoMigrate(&User{})
//
//}
//func insert1() {
//	user := User{
//		Name:     "tom",
//		Age:      20,
//		Birthday: time.Now(),
//		Active:   false,
//	}
//	db.Create(&user)
//}
//func del1() {
//
//	var user User
//	db.First(&user)
//	db.Delete(&user) //软删除
//}
//func del2() {
//	db.Delete(&User{}, 2)
//}
//func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
//	println("a")
//	return nil
//}
//func main() {
//	del2()
//}
