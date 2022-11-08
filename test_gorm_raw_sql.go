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
//func testRaw() {
//	type Result struct {
//		ID   int
//		Name string
//		Age  int
//	}
//	var result Result
//	db.Raw("select id,name,age from users where name=?", "tom").Scan(&result)
//	fmt.Printf("%v\n", result)
//}
//func testRaw2() {
//	db.Exec("Update users set age=? where name=?", 100, "tom")
//}
//func createTable() {
//	db.AutoMigrate(&User{})
//
//}
//func testraw3() {
//	var name string
//	var age int
//	//row := db.Table("users").Where("name=?", "tom").Select("name", "age").Row()
//	//row.Scan(&name, &age)
//	//fmt.Printf("name:%v,age:%v", name, age)
//
//	rows, err := db.Model(&User{}).Where("age> ?", 18).Select("name,age").Rows()
//	if err != nil {
//		panic("出错")
//	}
//
//	for rows.Next() {
//
//		rows.Scan(&name, &age)
//		fmt.Printf("name:%v,age:%v\n", name, age)
//	}
//}
//func main() {
//	testraw3()
//}
