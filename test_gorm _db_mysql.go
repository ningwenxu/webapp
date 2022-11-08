package main

//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"time"
//)
//
//var db *gorm.DB
//
//func createDB() {
//
//}
//func init() {
//	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
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
//}
//
//func createTable() {
//	db.AutoMigrate(&User{})
//}
//func insert1() {
//	user := User{
//		Name:     "tom",
//		Age:      20,
//		Birthday: time.Now(),
//	}
//	result := db.Create(&user)
//	fmt.Printf("result:%v\n", result.RowsAffected)
//	fmt.Printf("ID:%v\n", user.ID)
//}
//func insert2() {
//	user := User{
//		Name:     "tom",
//		Age:      20,
//		Birthday: time.Now(),
//	}
//	db.Select("Name", "Age", "CreatedAt").Create(&user)
//}
//func insert3() {
//	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
//	db.Create(&users)
//}
//func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//	fmt.Println("BeforeCreate...")
//	return
//}
//func insert4() {
//	db.Model(&User{}).Create(map[string]interface{}{
//		"Name": "jinzhu", "Age": 18,
//	})
//}
//func main() {
//	//createTable()
//	//insert1()
//	//insert2()
//	//insert3()
//	insert4()
//}
