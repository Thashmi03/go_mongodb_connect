package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct{
	Building string `json:"building" bson:"building,required"`
	Coord [] float64`json:"coord" bson:"coord,required"`
	Street string`json:"street" bson:"street,required"`
	Zipcode string`json:"zipcode" bson:"zipcode,required"`
}
type Grades struct{
	Date time.Time`json:"time" bson:"time,required"`
	Gra string`json:"grade" bson:"grade,required"`
	Score int`json:"score" bson:"score,required"`
}
type Res struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Add Address `json:"address" bson:"address,required"`
	Borough string `json:"borough" bson:"borough,required"`
	Cuisine string`json:"cuisine" bson:"cuisine,required"`
	Grade []Grades `json:"grade" bson:"grade,required"` 
	Name string `json:"name" bson:"name,required"`
	Res_id string `json:"res_id" bson:"res_id,required"`
}