package main

//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//var db *sql.DB
//
//func inittDB() (err error) {
//	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func insert(username string, password string) {
//	s := "insert into user_tbl (username,password) values(?,?)"
//	r, err := db.Exec(s, username, password)
//	if err != nil {
//		fmt.Printf("%v\n", err)
//	} else {
//		i, _ := r.LastInsertId()
//		fmt.Printf("%v\n", i)
//	}
//}
//
//type User struct {
//	id       int
//	username string
//	password string
//}
//
//func quaryonerow() {
//	s := "select * from user_tbl where id=?"
//	var u User
//	err := db.QueryRow(s, 2).Scan(&u.id, &u.username, &u.password)
//	if err != nil {
//		fmt.Printf("%v\n", err)
//	} else {
//		fmt.Printf("%v\n", u)
//	}
//}
//func quaryMayRow() {
//	s := "select * from user_tbl "
//	var u User
//	r, err := db.Query(s)
//	defer r.Close()
//	if err != nil {
//		fmt.Printf("%v\n", err)
//	} else {
//		for r.Next() {
//			r.Scan(&u.id, &u.username, &u.password)
//			fmt.Printf("%v\n", u)
//		}
//	}
//}
//func updateQuary() {
//	s := "update user_tbl set username=?,password=? where id=? "
//
//	r, err := db.Exec(s, "big kite", "31352", 2)
//
//	if err != nil {
//		fmt.Printf("%v\n", err)
//	} else {
//		i, _ := r.RowsAffected()
//		fmt.Printf("%v\n", i)
//	}
//
//}
//func delete() {
//	s := "delete from user_tbl  where id=? "
//
//	r, err := db.Exec(s, 1)
//
//	if err != nil {
//		fmt.Printf("%v\n", err)
//	} else {
//		i, _ := r.RowsAffected()
//		fmt.Printf("%v\n", i)
//	}
//}
//func main() {
//	//db, err := sql.Open("mysql", "root:root@/go_db")
//	//if err != nil {
//	//	panic(err)
//	//}
//	//// See "Important settings" section.
//	//db.SetConnMaxLifetime(time.Minute * 3)
//	//db.SetMaxOpenConns(10)
//	//db.SetMaxIdleConns(10)
//	//print(db)
//	err := inittDB()
//	if err != nil {
//		fmt.Printf("err:%v\n", err)
//	} else {
//		fmt.Printf("%v\n", "连接成功")
//	}
//	//quaryMayRow()
//	insert("aga", "asdq")
//}
