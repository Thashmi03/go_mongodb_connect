package main

import (
	"context"
	"fmt"
	"mongodb-dal/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main(){
	ctx:=context.TODO()
	//connect to mongodb
	mongoconn:=options.Client().ApplyURI(constants.ConnectionString)
	mongoclient,err:=mongo.Connect(ctx,mongoconn)
	if err!=nil{
		panic(err)
	}
	//readpref.Primary() asks whether it connects to first replica in cloud
	
	if err:=mongoclient.Ping(ctx,readpref.Primary());err!=nil{
		panic(err)
	}
	fmt.Println("MongoDB successfully connected...")

}