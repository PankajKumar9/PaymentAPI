package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id                        primitive.ObjectID `json:"_id" bson:"_id"`
	Info                      TransactionData    `json:"info" bson:"info"`
	CancellationInfo          TransactionData    `json:"cancellationInfo" bson:"cancellationInfo"`
	CancellationTransactionId string             `json:"cancellationTransactionId" bson:"cancellationTransactionId"`
	ParentTransactionId       string             `json:"parentTransactionId" bson:"parentTransactionId"`
	ParentInfo                TransactionData    `json:"parentInfo" bson:"parentInfo"`
}



type TransactionData struct {
	Account        string `json:"account" bson:"account"`
	Amount         string `json:"amount" bson:"amount"`
	Balance        string `json:"balance" bson:"balance"`
	Comments       string `json:"comments" bson:"comments"`
	Date           string `json:"date" bson:"date"`
	ImpactAmount   string `json:"impactAmount" bson:"impactAmount"`
	PaymentSummary string `json:"paymentSummary" bson:"paymentSummary"`
	PaymentNumber  string `json:"paymentNumber" bson:"paymentNumber"`
	ResultCode     string `json:"resultCode" bson:"resultCode"`
	Status         string `json:"Status" bson:"Status"`
}
