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
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	//it gets the data from database
	// bson.mfor greater than check
	filter := bson.M{"account_id":bson.D{{"$gt",7000}}}
	// result stores the data of 10 restaurant
	options := options.Find().SetSort(bson.D{{"transaction_count", 1}}).SetSkip(30).SetLimit(10)
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
// a>7000,tc==100
/*db.getCollection("transactions").find({"$and":[
    {"transaction_count":{"$eq":100}},
    {"account_id":{"$lte":700000}}
    ]})
*/
// sum
/*db.getCollection("transactions").aggregate([{
    $group:{
        _id:null,
        sum:{
            $sum:"$transaction_count"
        }
    }
}])*/

//inside trans amount>3000
// 
// sum of all transacations(amounts)
/*
db.getCollection("transactions").aggregate([{
    $project:{
        sum:{
            $sum:"$transactions.amount"
        }
    }
}])*/




func FetchAggregatedTransactions(){
	ctx,_:=context.WithTimeout(context.Background(),100*time.Second)
	matchStage:=bson.D{{"$match",bson.D{{"account_id",278866}}}}
	groupStage:=bson.D{
		{"$group",bson.D{
				{"_id","$account_id"},
				{"total_count",bson.D{{"$sum","$transaction_count"}}},
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