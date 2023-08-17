package services

import (
	"context"
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
	filter := bson.D{}
	// result stores the data of 10 restaurant
	options:=options.Find().SetSort(bson.D{{"transaction_count",-1}}).SetSkip(30).SetLimit(10)
	result, err := TransactionContext().Find(ctx, filter,options)
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
			transactions= append(transactions, res)
		}
		//checking for error
		if err := result.Err(); err != nil {
			return nil, err
		}
		return transactions, nil
	}
}
