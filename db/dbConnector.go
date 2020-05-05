package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mall/model"
	"time"
)

const (
	URI = "mongodb://106.15.239.125:27017/mall"
)

// 连接数据库
func Connect(dbName string) *mongo.Database {
	// 上下文
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// 创建连接
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	// 测试
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo has connected!")

	db := client.Database(dbName)
	return db
}

// 添加一个
func InsertOne(dbName, cName string, doc interface{}) error {
	db := Connect(dbName)
	defer db.Client().Disconnect(context.TODO())
	c := db.Collection(cName)
	_, err := c.InsertOne(context.TODO(), doc)
	return err
}

// 添加多个
func InsertMany(dbName, cName string, docs []interface{}) error {
	db := Connect(dbName)
	defer db.Client().Disconnect(context.TODO())

	c := db.Collection(cName)
	_, err := c.InsertMany(context.TODO(), docs)
	return err
}

// 查询
func FindOne(dbName, cName string, filter, result interface{}) error {
	db := Connect(dbName)
	defer db.Client().Disconnect(context.TODO())

	c := db.Collection(cName)
	err := c.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// 查询多个
func FindAll(dbName, cName string, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	db := Connect(dbName)
	defer db.Client().Disconnect(context.TODO())
	c := db.Collection(cName)

	var findoptions = &options.FindOptions{}
	findoptions.SetLimit(2)
	findoptions.SetSkip(0)
	findoptions.SetSort(map[string]int{"price": -1})

	cur, err := c.Find(context.TODO(), filter, findoptions)

	//err = cur.All(context.TODO(), results)
	//if err != nil {
	//	log.Fatal("cur.all:", err)
	//}

	for cur.Next(context.TODO()) {
		var o model.OrderItem
		err := cur.Decode(&o)
		if err!=nil {
			log.Fatal(err)
		}
		fmt.Println(&o)
	}

	defer cur.Close(context.TODO())
	return err
}

// 更新
func Update(dbName, cName string, query, update interface{}) error {
	db := Connect(dbName)
	defer db.Client().Disconnect(context.TODO())
	c := db.Collection(cName)

	options := &options.UpdateOptions{}
	// 有则更新无则添加
	options.SetUpsert(true)

	_, err := c.UpdateOne(context.Background(), query, update, options)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func UpdateAll(dbName, cName string, query, update interface{}) error {
	db := Connect(dbName)
	defer db.Client().Disconnect(context.TODO())
	c := db.Collection(cName)

	options := &options.UpdateOptions{}
	// 有则更新无则添加
	options.SetUpsert(true)

	_, err := c.UpdateMany(context.Background(), query, update, options)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func Remove(dbName, cName string, query interface{}) error {
	db := Connect(dbName)
	defer db.Client().Disconnect(context.TODO())
	c := db.Collection(cName)
	_, err := c.DeleteOne(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// 删除
func RemoveAll(dbName, cName string, query interface{}) error {
	db := Connect(dbName)
	defer db.Client().Disconnect(context.TODO())
	c := db.Collection(cName)
	_, err := c.DeleteMany(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
