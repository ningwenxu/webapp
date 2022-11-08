package main

//
//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//)
//
//type Student struct {
//	Name string
//	Age  int
//}
////
//
//var client *mongo.Client
//
//func initDB() {
//	// 设置客户端连接配置
//	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//
//	// 连接到MongoDB
//	var err error
//	c, err := mongo.Connect(context.TODO(), clientOptions) //返回数据库连接
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 检查连接
//	err = c.Ping(context.TODO(), nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("MongoDB 连接成功")
//	client = c
//}
//func insert() {
//	initDB()
//	s1 := Student{
//		Name: "tom",
//		Age:  20,
//	}
//	collection := client.Database("golang_db").Collection("Student")
//	ior, err := collection.InsertOne(context.TODO(), s1)
//	if err != nil {
//		fmt.Printf("err:%v\n", err)
//	} else {
//		fmt.Printf("ior.InsertID:%v\n", ior.InsertedID)
//	}
//}
//func insertMany() {
//	initDB()
//	c := client.Database("golang_db").Collection("Student")
//	s1 := Student{
//		Name: "kite",
//		Age:  20,
//	}
//	s2 := Student{
//		Name: "rose",
//		Age:  20,
//	}
//	stus := []interface{}{s1, s2}
//	imr, err := c.InsertMany(context.TODO(), stus)
//	if err != nil {
//		fmt.Printf("err:%v\n", err)
//	} else {
//		fmt.Printf("ior.InsertID:%v\n", imr.InsertedIDs)
//	}
//}
//
//func find() {
//	defer client.Disconnect(context.TODO())
//	c := client.Database("golang_db").Collection("Student")
//	c2, err := c.Find(context.TODO(), bson.D{{"name", "tom"}}) //查找姓名是tom的
//	//c2, err := c.Find(context.TODO(), bson.D{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer c2.Close(context.TODO())
//	for c2.Next(context.TODO()) {
//		var result bson.D
//		c2.Decode(&result)
//
//		fmt.Printf("result:%v\n", result)
//		fmt.Printf("result:%v\n", result.Map())
//	}
//}
//
//func update() {
//	ctx := context.TODO()
//	c := client.Database("golang_db").Collection("Student")
//	update := bson.D{{"$set", bson.D{{"name", "big kite"}, {"Age", 22}}}}
//	ur, err := c.UpdateMany(ctx, bson.D{{"name", "kite"}}, update)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("modifiedcount:%v\n", ur.ModifiedCount)
//}
//
//func delete() {
//	ctx := context.TODO()
//	c := client.Database("golang_db").Collection("Student")
//
//	ur, err := c.DeleteMany(ctx, bson.D{{"name", "big kite"}})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("modifiedcount:%v\n", ur.DeletedCount)
//}
//func main() {
//	//initDB()
//	//insert()
//	//insertMany()
//	initDB()
//	//update()
//	delete()
//}
//
////func main() {
////	////filter 过滤器  sql where
////	//d := bson.D{{"name", "tom"}}
////	//print(d)
////
////}
