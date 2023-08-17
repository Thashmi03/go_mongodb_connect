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
func ProductCont()*mongo.Collection{
	// it checks for the connection
	client,_:=config.ConnectDatabase()
	//returing the collection
	return config.GetCollection(client,"sample_restaurants","restaurants")
}

func FindRes() ([]*models.Res, error) {
	// this is used for deleteing after the given time
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//it gets the data from database
	filter := bson.D{}
	// result stores the data of 10 restaurant
	result, err := ProductCont().Find(ctx, filter, options.Find().SetLimit(10))
	// error checking
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		// creating array to store the elements
		var rest []*models.Res
		for result.Next(ctx) {
			post := &models.Res{}  
			err := result.Decode(post)
			if err != nil {
				return nil, err
			}
			// adding the restaurant elements into array
			rest = append(rest, post)
		}
		//checking for error
		if err := result.Err(); err != nil {
			return nil, err
		}
		return rest, nil
	}

}
