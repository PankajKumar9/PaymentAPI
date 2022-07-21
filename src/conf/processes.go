package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/PankajKumar9/PaymentAPI/src/conf/constants"
	"github.com/PankajKumar9/PaymentAPI/src/database"
	"github.com/PankajKumar9/PaymentAPI/src/models"
	"github.com/PankajKumar9/PaymentAPI/src/utility"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUsers() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	NewUsers, Err := ioutil.ReadFile("src/signupform_sample.json")
	if Err != nil {
		fmt.Println("Yahi fatt gya ye toh ")
	}
	log.Println(utility.Info(fmt.Sprintf(string(NewUsers))))
	users := []models.User{}
	// cms contract from grpc

	err2 := json.Unmarshal([]byte(NewUsers), &users)
	log.Println(utility.Info(fmt.Sprintf("%v", err2)))

	for _, u := range users {
		database.InsertUser(u)
	}
}

func process(email string, amount float64, mode string, User models.User) (error, string) {
	if mode == constants.DEBIT && User.Balance < amount {
		return errors.New(constants.INSUFFICIENT_FUNDS), ""
	}
	t := models.Transaction{}
	t.Id = database.InsertTransaction(t)
	t.Info.Account = User.AccountNumber
	t.Info.Amount = amount
	if mode == constants.CREDIT {
		User.Balance += amount
	} else {
		User.Balance -= amount
	}
	t.Info.Date = time.Now().GoString()
	t.Info.User = User
	t.Info.Balance = User.Balance
	t.Info.Status = mode
	database.UpdateTransaction(t)
	User.History = append(User.History, t.Id)
	database.UpdateUser(User)
	return nil, t.Id.Hex()
}

//if flag false then dont move the money
//TODO implement some sort of read-write lock
func CancelTransaction(ID string, flag bool) {
	Id, _ := primitive.ObjectIDFromHex(ID)
	t, _, _ := database.FindTransaction(Id)
	User, _, _ := database.FindUser(t.Info.User.Email)
	mode := t.Info.Status
	transactionId := ""
	if mode == constants.CREDIT {
		_, transactionId = process(User.Email, t.Info.Amount, constants.DEBIT, User)
	} else {
		_, transactionId = process(User.Email, t.Info.Amount, constants.CREDIT, User)
	}
	t.CancellationTransactionId = transactionId
	database.UpdateTransaction(t)
	//TODO make this dynamic with process function itself
	TransId, _ := primitive.ObjectIDFromHex(transactionId)
	t2, _, _ := database.FindTransaction(TransId)
	t2.ParentTransactionId = t.Id.Hex()
	database.UpdateTransaction(t2)
}
