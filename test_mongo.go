package main

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
//
//func main() {
//	initDB()
//}
