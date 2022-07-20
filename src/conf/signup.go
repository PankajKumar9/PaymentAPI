package conf

import (
	"encoding/json"
	"errors"

	"github.com/PankajKumar9/PaymentAPI/src/database"
	"github.com/PankajKumar9/PaymentAPI/src/models"

	routing "github.com/qiangxue/fasthttp-routing"
)

func Signup(signupform string, c *routing.Context) error {

	newUser := &models.User{}
	err := json.Unmarshal([]byte(signupform), newUser)
	if err != nil {
		return err
	}

	pass, msg := validateUser(c)
	if !pass {
		return errors.New(msg)
	}
	//todo check password == confirm password

	//TODO will send password for hashing , wont write the password

	database.InsertUser(*newUser)
	return nil
}
func validateUser(c *routing.Context) (bool, string) {

	email := string(c.Request.Header.Peek("email"))
	if len(email) == 0 {
		return false, "NO EMAIL SENT"
	}
	_, pass, _ := database.FindUser(email)
	if pass {
		return false, "User already exists"
	}

	//TODO : write logic to check mandetory fields
	return true, ""
}
