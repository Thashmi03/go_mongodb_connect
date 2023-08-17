package main

import (
	"fmt"
	"mongodb-dal/services"
	// "mongodb-dal/config"
	// "mongodb-dal/models"
	// "mongodb-dal/services"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var(
	mongoclient *mongo.Client
)
func main(){
	// client,_:=config.ConnectDatabase()
	// config.GetCollection(client,"sample training")
	
	
	// inserting one product
	// product:=models.Product{ID: primitive.NewObjectID(),Name: "OnePlus",Price: 14000,Description: "Good"}
	
	// inserting multiple products
	// product:=[]interface{}{models.Product{ID: primitive.NewObjectID(),Name: "OnePlus",Price: 14000,Description: "Good"},
	// models.Product{ID: primitive.NewObjectID(),Name: "Redmi9",Price: 9000,Description: "avg"},
	// models.Product{ID: primitive.NewObjectID(),Name: "Nokia",Price: 11000,Description: "Good"}}
	// services.InsertProductList(product)
	// p1:=models.Product{ID: primitive.NewObjectID(),Name: "realme",Price: 11000,Description: "Good"}
	// services.InsertProduct(p1)
	// services.FindProducts()


	fmt.Println("MongoDB successfully connected...")
	// pro,_:=services.FindTransactions()
	// for _,p:=range pro{
	// 	// fmt.Printf("{\nid:%v,\naccount_id:%v,\ntransaction_count:%v,\nbucket_start_date:%v,\nbucket_end_date:%v\n}\n",p.ID,p.AccountId,p.Trans_count,p.Bucket_start,p.Bucket_end)
	// 	fmt.Println(p)
	// }
	// services.FetchAggregatedTransactions()
	
	transaction,_:=services.UpdateTransaction(4431789,637483)
	fmt.Println(transaction.ModifiedCount)
	
}