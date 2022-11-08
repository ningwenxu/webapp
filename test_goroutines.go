//go:build linux
// +build linux

package main

import (
	"encoding/json"
	"fmt"
)

//	func createFile() {
//		f, err := os.Create("a.txt")
//		if err != nil {
//			fmt.Printf("err:%v\n", err)
//
//		} else {
//			fmt.Printf("f.name:%v\n", f.Name())
//		}
//	}
//
//	func createGir() {
//		err := os.Mkdir("a", os.ModeAppend)
//		if err != nil {
//			fmt.Printf("err:%v\n", err)
//
//		}
//	}
//
//	func testReadDir() {
//		fi, _ := ioutil.ReadDir(".")
//		for _, v := range fi {
//			fmt.Printf("v.Name():%v\n", v.Name())
//		}
//	}
type Person struct {
	Name  string
	Age   int
	Email string
}

func test1() {
	p := Person{
		Name:  "Tom",
		Age:   21,
		Email: "2389129",
	}
	b, _ := json.Marshal(p)
	fmt.Printf("p:%v\n", string(b))
}
func test2() {
	b := []byte(`{"Name":"Tom","Age":21,"Email":"2389129"}`)
	var p Person
	json.Unmarshal(b, &p)
	fmt.Printf("p:%v\n", p)
}
func main() {
	test2()
}
