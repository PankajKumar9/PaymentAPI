package database

import (
	"context"
	"fmt"
	"log"

	"github.com/PankajKumar9/PaymentAPI/src/models"
	"github.com/PankajKumar9/PaymentAPI/src/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

const dbName = "PaymentAPI"
const colUsers = "users"
const colTransactions = "transaction"

var CollectionUsers *mongo.Collection
var CollectionTransactions *mongo.Collection

func init() {
	//client options

	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	CollectionUsers = client.Database(dbName).Collection(colUsers)
	CollectionTransactions = client.Database(dbName).Collection(colTransactions)
	fmt.Println("Collection instance is ready")

}
func GetCollectionUsers() *mongo.Collection {
	return CollectionUsers
}
func GetCollectionTransactions() *mongo.Collection {
	return CollectionTransactions
}
func InsertUser(user models.User) {
	inserted, err := CollectionUsers.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted user in db with id", inserted.InsertedID)

}
func UpdateUser(user models.User) {

	filter := bson.M{"_id": user.Id}
	updatedUser, err := bson.Marshal(user)
	if err != nil {
		log.Println(utility.Info("could not marshal user to bson"))
	}
	UserPrimitive := primitive.D{}
	err = bson.Unmarshal([]byte(updatedUser), &UserPrimitive)
	update := bson.M{"$set": UserPrimitive}

	_, err = CollectionUsers.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(utility.Info("could not update user in mongo"))
	}

}

func InsertTransaction(transaction models.Transaction) {
	inserted, err := CollectionTransactions.InsertOne(context.Background(), transaction)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted Transaction in db with id", inserted.InsertedID)

}
func UpdateTransaction(transaction models.Transaction) {

	filter := bson.M{"_id": transaction.Id}
	updatedTransaction, err := bson.Marshal(transaction)
	if err != nil {
		log.Println(utility.Info("could not marshal transaction to bson"))
	}
	TransactionPrimitive := primitive.D{}
	err = bson.Unmarshal([]byte(updatedTransaction), &TransactionPrimitive)
	update := bson.M{"$set": TransactionPrimitive}

	_, err = CollectionTransactions.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(utility.Info("could not update transactions in mongo"))
	}

}
