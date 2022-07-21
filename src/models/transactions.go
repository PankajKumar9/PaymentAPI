package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id"`
	Info TransactionData    `json:"info" bson:"info"`

	CancellationTransactionId string `json:"cancellationTransactionId" bson:"cancellationTransactionId"`
	ParentTransactionId       string `json:"parentTransactionId" bson:"parentTransactionId"`

	InverseTranactionId string `json:"inverseTransactionId" bson:"inverseTransactionId"`
}
type CreditRequest struct {
	Email    string  `json:"email " bson:"email "`
	Password string  `json:"password" bson:"password"`
	Amount   float64 `json:"amount" bson:"amount"`
}

type TransactionData struct {
	User           User    `json:"user" bson:"user"`
	Account        string  `json:"account" bson:"account"`
	Amount         float64 `json:"amount" bson:"amount"`
	Balance        float64 `json:"balance" bson:"balance"`
	Comments       string  `json:"comments" bson:"comments"`
	Date           string  `json:"date" bson:"date"`
	ImpactAmount   string  `json:"impactAmount" bson:"impactAmount"`
	PaymentSummary string  `json:"paymentSummary" bson:"paymentSummary"`
	PaymentNumber  string  `json:"paymentNumber" bson:"paymentNumber"`
	ResultCode     string  `json:"resultCode" bson:"resultCode"`
	Status         string  `json:"Status" bson:"Status"`
}

// Status => Debit/Credit
type SendRequest struct {
	From struct {
		Email    string  `json:"email " bson:"email "`
		Password string  `json:"password" bson:"password"`
		Amount   float64 `json:"amount" bson:"amount"`
	} `json:"from " bson:"from "`
	To struct {
		Email string `json:"email " bson:"email "`
		ID    string `json:"Id" bson:"Id"`
	} `json:"to" bson:"to"`
}
type RefundResponse struct {
	Email         string `json:"email " bson:"email "`
	Password      string `json:"password" bson:"password"`
	TransactionID string `json:"transactionId" bson:"transactionId"`
}
