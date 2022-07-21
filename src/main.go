package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/PankajKumar9/PaymentAPI/src/RestAPIs"
	"github.com/PankajKumar9/PaymentAPI/src/database"
	"github.com/PankajKumar9/PaymentAPI/src/testing"
	"github.com/PankajKumar9/PaymentAPI/src/utility"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"
)

func main() {
	app := &cli.App{
		Name: "PaymentAPIT",
		Action: func(*cli.Context) error {
			drive()
			return nil
		},
	}
	fmt.Println("Check arguments")
	fmt.Println("%v", os.Args)
	fmt.Println("---------")
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

var Router *routing.Router

func drive() {

	//test logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println(utility.Info("setting up configurations"))
	database.Init()
	Router = routing.New()
	RestAPIs.InitApis(Router)

	//
	log.Println(utility.Info("003"))
	var wg sync.WaitGroup
	go HttpListen(Router)
	wg.Add(1)
	log.Println(utility.Info("004"))
	testing.TestRun()

}
func HttpListen(Router *routing.Router) {
	log.Println(utility.Info("001"))
	err := fasthttp.ListenAndServe(":3001", Router.HandleRequest)
	if err != nil {
		log.Println(utility.Info(err))
	}

}

// user1 := models.User{}
// user2 := models.User{}
// t1 := models.Transaction{}
// t2 := models.Transaction{}

// database.InsertTransaction(t1)
// database.InsertTransaction(t2)
// database.InsertUser(user1)
// database.InsertUser(user2)
// user1 := models.User{}
// y, _ := json.Marshal(user1)
// x := string(y)
// log.Println(utility.Info(x))
