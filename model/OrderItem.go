package model

import "time"

type OrderItem struct {
	Price float64 `bson:"price"`
	Count int `bson:"count"`
	Time time.Time `bson:"time"`
	Title string `bson:"title"`
	Desc string `bson:"desc"`
}