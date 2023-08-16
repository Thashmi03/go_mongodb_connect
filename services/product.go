package services

import (
	"context"
	"fmt"
	"mongodb-dal/config"
	"mongodb-dal/models"
	"time"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
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
	return config.GetCollection(client,"inventory","product")

}
func InsertProduct(product models.Product){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	result,err:=ProductContext().InsertOne(ctx,product)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("result",result)
	
}

func InsertProductList(products []interface{}){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	result,err:=ProductContext().InsertMany(ctx,products)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("result",result)
	
}

//finding products
func FindProducts()([]*models.Product,error){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	filter:=bson.D{{"price",11000}}
	result,err := ProductContext().Find(ctx,filter)
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}else{
		//build the array of products
		var products[]*models.Product
		for result.Next(ctx){
			post:=&models.Product{}
			err:=result.Decode(post)
			if err!=nil{
				return nil,err
			}
			
			products=append(products, post)
		}
		if err:=result.Err();err!=nil{
			return nil,err
		}
		if len(products)==0{
			return []*models.Product{},nil
		}
		return products,nil
	}
}


