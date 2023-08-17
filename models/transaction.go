package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Transaction struct{
	ID 				primitive.ObjectID `json:"_id" bson:"_id"`
	AccountId 		int `json:"account_id" bson:"account_id"`
	Trans_count 	int `json:"trans_count" bson:"transaction_count"`
	Bucket_start 	primitive.DateTime `json:"bucket_start_date" bson:"bucket_start_date"`
	Bucket_end 		primitive.DateTime`json:"bucket_end_date" bson:"bucket_end_date"`
	Trans 			[] Inner_Transaction `json:"transactions" bson:"transactions"`
}
type Inner_Transaction struct{
	Date 			primitive.DateTime `json:"date" bson:"date"`
	Amount 			int `json:"amount" bson:"amount"`
	Trans_code 		string`json:"transaction_code" bson:"transaction_code"`
	Symbol 			string`json:"symbol" bson:"symbol"`
	Price 			string`json:"price" bson:"price"`
	Total 			string`json:"total" bson:"total"`
}
