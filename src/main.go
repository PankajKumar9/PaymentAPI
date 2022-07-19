package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PankajKumar9/PaymentAPI/src/utility"
	"github.com/urfave/cli/v2"
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
func drive() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println(utility.Info("something"))
}
