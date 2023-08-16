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

// "context"
// "fmt"
// "mongodb-dal/constants"

// "go.mongodb.org/mongo-driver/mongo"
// "go.mongodb.org/mongo-driver/mongo/options"
// "go.mongodb.org/mongo-driver/mongo/readpref"
var(
	mongoclient *mongo.Client
)
func main(){
	// client,_:=config.ConnectDatabase()
	// config.GetCollection(client,"sample training")
	
	fmt.Println("MongoDB successfully connected...")
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


	pro,_:=services.FindRes()
	for i,p:=range pro{
		fmt.Println(i+1,"Name:",p.Name,"-Cuisine",p.Cuisine)
	}
	
	
}