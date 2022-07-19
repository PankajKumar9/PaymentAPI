package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id            primitive.ObjectID `json:"_id" bson:"_id"`
	Name          string             `json:"name" bson:"name"`
	Balance       float64            `json:"balance" bson:"balance"`
	Type          string             `json:"type" bson:"type"`
	AccountType   string             `json:"accountType" bson:"accountType"`
	Status        string             `json:"status" bson:"status"`
	BankName      string             `json:"bank_name" bson:"bank_name"`
	Country       string             `json:"country" bson:"country"`
	Currency      string             `json:"currency" bson:"currency"`
	Customer      interface{}        `json:"customer" bson:"customer"`
	Fingerprint   string             `json:"fingerprint" bson:"fingerprint"`
	Last4         string             `json:"last4" bson:"last4"`
	History       Transaction        `json:"history" bson:"history"`
	RoutingNumber string             `json:"routing_number" bson:"routing_number"`
}
