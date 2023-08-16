package services

import (
	"context"
	"fmt"
	"mongodb-dal/config"
	"mongodb-dal/models"
	"time"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



// 	code for insert one
// *************************
// func InsertProduct(){
// 	var product models.Product
// 	product.ID=primitive.NewObjectID()
// 	product.Name ="iPhone"
// 	product.Price=115000
// 	product.Description="It is an awesome phone"
// 	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
// 	client,_:=config.ConnectDatabase()
// 	var productCollection *mongo.Collection=config.GetCollection(client,"inventory","products")
// 	result,err:=productCollection.InsertOne(ctx,product)
// 	if err!=nil{
// 		fmt.Println(err)
// 	}
// 	fmt.Println("result",result)
	
	
// }

// code for insert with function
// **********************************
func ProductContext()*mongo.Collection{
	client,_:=config.ConnectDatabase()
	return config.GetCollection(client,"inventory","products")

}
func InsertProduct(product models.Product)(*mongo.InsertOneResult,error){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	result,err:=ProductContext().InsertOne(ctx,product)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("result",result)
	return result,nil
}

func InsertProductList(products []interface{})(*mongo.InsertManyResult,error){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	result,err:=ProductContext().InsertMany(ctx,products)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("result",result)
	return result,nil
}

