package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mall/db"
	"mall/model"
	"time"
)

func main() {
	var results []model.OrderItem

	var findoptions = &options.FindOptions{}
	findoptions.SetLimit(5)
	findoptions.SetSkip(0)
	findoptions.SetSort(map[string]int{"price": -1})
	//
	//
	err := db.FindAll("mall", "orderItem", bson.M{"price": bson.M{"$lt": 100 }}, &results, findoptions)
	////
	if err == nil {
		//fmt.Println("result", results)
		for _, v := range results {
			fmt.Println("row:", v.Price)
		}
	}

	// remove
	//err := db.Remove("mall", "orderItem", bson.M{"price": 68.8})


	// 更新
	//err := db.Update("mall", "orderItem", bson.M{"price": 8.8}, bson.M{"$set":bson.M{"price": 88.8}})

	if err!=nil {
		fmt.Println(err)
	}
}

func insertItem () {
	items := [] model.OrderItem{
		{1, 1,time.Now(), "洗衣机", ""},
		{2, 2,time.Now(), "冰箱", ""},
		{5, 3,time.Now(), "哈密瓜", ""},
		{8, 1,time.Now(), "狄安娜", ""},
		{9.8, 1,time.Now(), "笔记本", ""},
		{10.8, 1,time.Now(), "美女", ""},
	}

	var _items []interface{}
	for _, v := range items {
		_items = append(_items, v)
	}
	db.InsertMany("mall", "orderItem", _items)
}