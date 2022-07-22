package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id            primitive.ObjectID `json:"_id" bson:"_id"`
	Name          string             `json:"name" bson:"name"`
	Email         string             `json:"email" bson:"email"`
	Balance       float64            `json:"balance" bson:"balance"`
	Type          string             `json:"type" bson:"type"`
	AccountType   string             `json:"accountType" bson:"accountType"`
	AccountNumber string             `json:"accountNumber" bson:"accountNumber"`
	Status        string             `json:"status" bson:"status"`
	BankName      string             `json:"bankName" bson:"bankName"`
	Country       string             `json:"country" bson:"country"`
	Currency      string             `json:"currency" bson:"currency"`

	Last4   []Transaction        `json:"last4" bson:"last4"`
	History []primitive.ObjectID `json:"history" bson:"history"`
	//TODO : remove the password field after testing
	//only password hash is in user properties
	Password     string `json:"password" bson:"password"`
	PasswordHash string `json:"passwordhash" bson:"passwordhash"`
}
