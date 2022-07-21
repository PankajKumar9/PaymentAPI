package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/PankajKumar9/PaymentAPI/src/conf"
	"github.com/PankajKumar9/PaymentAPI/src/database"
	"github.com/PankajKumar9/PaymentAPI/src/models"
	"github.com/PankajKumar9/PaymentAPI/src/utility"
)

type Test struct {
	U         models.User
	ReqCredit models.CreditRequest
	ReqDebit  models.CreditRequest
	ReqSend   models.SendRequest
	ReqRefund models.RefundResponse
}

func TestRun() {
	log.Println(utility.Info("002"))
	conf.CreateUsers()
	StressTesting()
}
func StressTesting() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var test [10]Test
	FillUsers(&test)
	log.Println(utility.Info(test))
	for i, t := range test {
		test[i].ReqCredit.Amount = float64(rand.Intn(200))
		test[i].ReqCredit.Email = t.U.Email
		test[i].ReqCredit.Password = t.U.Password

		test[i].ReqDebit.Amount = float64(rand.Intn(200))
		test[i].ReqDebit.Email = t.U.Email
		test[i].ReqDebit.Password = t.U.Password

		test[i].ReqSend.From.Amount = float64(rand.Intn(200))
		test[i].ReqSend.From.Email = t.U.Email
		test[i].ReqSend.From.Password = t.U.Password
		if i < 5 {
			test[i].ReqSend.To.Email = test[5+rand.Intn(5)].U.Email
		} else {
			test[i].ReqSend.To.Email = test[rand.Intn(5)].U.Email
		}
	}
	for _, t := range test {
		Lik := "http://localhost:3001/api/credit"
		x := t.ReqCredit
		log.Println(utility.Info(x))
		x.Amount = 100000
		S, _ := json.Marshal(x)
		by := bytes.NewBuffer(S)
		Postcalls(Lik, by)
	}
	for _, t := range test {
		Lik := "http://localhost:3001/api/debit"
		x := t.ReqDebit
		log.Println(utility.Info(x))
		x.Amount = float64(rand.Intn(200))
		S, _ := json.Marshal(x)
		by := bytes.NewBuffer(S)
		Postcalls(Lik, by)
	}
	for i, t := range test {
		if i == 5 {
			break
		}
		Lik := "http://localhost:3001/api/send"
		x := t.ReqSend
		log.Println(utility.Info(x))

		S, _ := json.Marshal(x)
		by := bytes.NewBuffer(S)
		Postcalls(Lik, by)
	}
	x, _, _ := database.GetTransactionsDataForTesting()
	log.Println(utility.Info(x))
	for _, transaction := range x {
		req := models.RefundResponse{}
		req.Email = transaction.Info.To
		req.Password = GetPassword(req.Email)
		req.TransactionID = transaction.Id.Hex()
		log.Println(utility.Info(req))
		Lik := "http://localhost:3001/api/refund"
		x := req
		log.Println(utility.Info(x))

		S, _ := json.Marshal(x)
		by := bytes.NewBuffer(S)
		Postcalls(Lik, by)

	}

}
func FillUsers(test *[10]Test) {

	//
	NewUsers, Err := ioutil.ReadFile("src/signupform_sample.json")
	if Err != nil {
		fmt.Println("Yahi fatt gya ye toh ")
	}

	users := []models.User{}
	// cms contract from grpc

	err2 := json.Unmarshal([]byte(NewUsers), &users)
	if err2 != nil {
		fmt.Println(err2)
	}
	// log.Println(utility.Info(users))

	//
	// fmt.Println(len(test))
	// fmt.Println(len(users))
	// log.Println(utility.Info(users[0]))
	// log.Println(utility.Info(test[0]))
	for i, _ := range test {
		//log.Println(i)
		test[i] = Test{}
		test[i].U = users[i]
		//log.Println(t.U)
		//log.Println(users[i])
	}
	//log.Println(utility.Info(test))
}
func testPostreq(test []Test) {
	fmt.Println("Chaliye shuru krte h")
	Link := "http://localhost:3001/api/credit"

	x := models.CreditRequest{}
	x.Email = "aaaa@bbbb.com"
	x.Password = "apples@pear"
	x.Amount = 1000
	S, _ := json.Marshal(x)
	resp, err := http.Post(Link, "application/json; charset=utf-8", bytes.NewBuffer(S))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Line number 29 ")
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func Postcalls(Lik string, by *bytes.Buffer) {
	resp, err := http.Post(Lik, "application/json; charset=utf-8", by)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Line number 180 ")
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
func GetPassword(email string) string {
	mapData := map[string]string{

		"aaaj@bbbk.com": "peach@melon",
		"aaai@bbbj.com": "melon@peach",
		"aaah@bbbi.com": "lime@kiwi",
		"aaag@bbbh.com": "kiwi@lime",
		"aaaf@bbbg.com": "plum@mango",
	}
	password, exists := mapData[email]
	if !exists {
		log.Println("YAHA TEST F@T@")
		return "NOT FOUND"
	}
	return password
}
