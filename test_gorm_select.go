//go:build linux
// +build linux

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
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

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time //要大写
}

func select1() {
	var user User
	//db.First(&user)
	//fmt.Printf("ID:%v\n", user.ID)
	db.Last(&user)
	fmt.Printf("ID:%v\n", user.ID)
}
func select2() {
	var users []User
	db.Find(&users, []int{1, 2, 3})

	for _, user := range users {
		fmt.Printf("ID:%v\n", user.ID)
	}
}
func select3() {
	var users []User
	result := db.Find(&users)
	fmt.Printf("count:%v\n", result.RowsAffected)
}
func main() {
	//select1()
	//select2()
	select3()
}
