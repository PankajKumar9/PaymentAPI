package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/PankajKumar9/PaymentAPI/src/RestAPIs"
	"github.com/PankajKumar9/PaymentAPI/src/database"
	"github.com/PankajKumar9/PaymentAPI/src/models"
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
	x := []models.User{}
	x = append(x, models.User{})
	x = append(x, models.User{})
	x = append(x, models.User{})
	x = append(x, models.User{})
	x = append(x, models.User{})
	y, _ := json.Marshal(x)
	a := string(y)
	log.Println(utility.Info(a))

	//test logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(utility.Info("something"))
	//set up router
	Router = routing.New()
	RestAPIs.InitApis(Router)

	// user1 := models.User{}
	// y, _ := json.Marshal(user1)
	// x := string(y)
	// log.Println(utility.Info(x))
	database.Init()

	fasthttp.ListenAndServe(":3001", Router.HandleRequest)
}

// user1 := models.User{}
// user2 := models.User{}
// t1 := models.Transaction{}
// t2 := models.Transaction{}

// database.InsertTransaction(t1)
// database.InsertTransaction(t2)
// database.InsertUser(user1)
// database.InsertUser(user2)
