package services

import (
	"context"
	"encoding/json"
	"fmt"
	"mongodb-dal/config"
	"mongodb-dal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TransactionContext() *mongo.Collection {
	client, _ := config.ConnectDatabase()
	return config.GetCollection(client, "sample_analytics", "transactions")
}
func FindTransactions() ([]*models.Transaction, error) {
	// this is used for deleteing after the given time
	ctx,cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	//it gets the data from database
	// bson.mfor greater than check
	filter := bson.M{"account_id":bson.D{{Key: "$gt",Value: 7000}}}
	// result stores the data of 10 restaurant
	options := options.Find().SetSort(bson.D{{Key: "transaction_count", Value: 1}}).SetSkip(30).SetLimit(10)
	result, err := TransactionContext().Find(ctx, filter, options)
	// error checking
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		// creating array to store the elements
		var transactions []*models.Transaction
		for result.Next(ctx) {
			res := &models.Transaction{}
			//unmarshal the current document
			err := result.Decode(res)
			if err != nil {
				return nil, err
			}
			// adding the restaurant elements into array
			transactions = append(transactions, res)
		}
		//checking for error
		if err := result.Err(); err != nil {
			return nil, err
		}
		return transactions, nil
	}
}


func FetchAggregatedTransactions(){
	ctx,cancel:=context.WithTimeout(context.Background(),100*time.Second)
	defer cancel()
	matchStage:=bson.D{{Key: "$match",Value: bson.D{{Key: "account_id",Value: 278866}}}}
	groupStage:=bson.D{
		{Key: "$group",Value: bson.D{
				{Key: "_id",Value: "$account_id"},
				{Key: "total_count",Value: bson.D{{Key: "$sum",Value: "$transaction_count"}}},
			}}}
	result,err:=TransactionContext().Aggregate(ctx,mongo.Pipeline{matchStage,groupStage})
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		var showWithInfo[]bson.M
		if err=result.All(ctx,&showWithInfo);err!=nil{
			panic(err)
		}
		formatted_data,err:=json.MarshalIndent(showWithInfo,""," ")
		if err!=nil{
			fmt.Println(err.Error())
		}else{
			fmt.Println(string(formatted_data))
		}
		
	}

}

func UpdateTransaction(intialValue int,newValue int)(*mongo.UpdateResult,error){
	ctx,cancel:=context.WithTimeout(context.Background(),100*time.Second)
	defer cancel()
	filter:=bson.D{{Key: "account_id",Value: intialValue}}
	update:=bson.D{{Key: "$set",Value: bson.D{{Key: "account_id",Value: newValue}}}}
	result,err:=TransactionContext().UpdateOne(ctx,filter,update)
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return result,nil
}
