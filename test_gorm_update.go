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
	Active   bool
}

func createTable() {
	db.AutoMigrate(&User{})
}
func insert() {
	user := User{
		Name:     "tom",
		Age:      20,
		Birthday: time.Now(),
		Active:   true,
	}
	db.Create(&user)
}
func update1() {
	var user User
	db.First(&user) ///id =1
	user.Name = "big tom"
	user.Age = 100
	db.Save(&user)
}
func update2() {
	db.Model(&User{}).Where("active=?", true).Update("name", "hello")
}
func update3() {
	var user User
	db.First(&user) ///id =1

	//db.Model(&user).Updates(User{Name: "helrrrlo", Age: 18, Active: false})
	db.Model(&user).Updates(map[string]interface{}{"Name": "helrrrlo", "Age": 18, "Active": false})
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeUpdate")
	return nil
}
func main() {
	update3()
}
