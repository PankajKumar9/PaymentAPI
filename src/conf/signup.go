package conf

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/PankajKumar9/PaymentAPI/src/conf/constants"
	"github.com/PankajKumar9/PaymentAPI/src/database"
	"github.com/PankajKumar9/PaymentAPI/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	routing "github.com/qiangxue/fasthttp-routing"
)

func Signup(signupform string, c *routing.Context) error {

	newUser := &models.User{}
	err := json.Unmarshal([]byte(signupform), newUser)
	if err != nil {
		return err
	}

	pass, msg := ValidateUser(*newUser)
	if !pass {
		return errors.New(msg)
	}
	//todo check password == confirm password

	//TODO will send password for hashing , wont write the password

	database.InsertUser(*newUser)
	return nil
}

func Credit(reqform string, c *routing.Context) error {
	req := &models.CreditRequest{}
	err := json.Unmarshal([]byte(reqform), req)
	if err != nil {
		return err
	}
	pass, msg := ValidateRequest(*req)
	if !pass {
		return errors.New(msg)
	}
	User, pass, _ := database.FindUser(req.Email)
	//TODO will deal with passwordHash and not the actual passwords
	if req.Password != User.Password {
		//TODO this static string has to be in constants package
		return errors.New("INCORRECT PASSWORD")
	}
	err, _ = process(req.Email, req.Amount, constants.CREDIT, User)
	return err

}
func Debit(reqform string, c *routing.Context) error {
	req := &models.CreditRequest{}
	err := json.Unmarshal([]byte(reqform), req)
	if err != nil {
		return err
	}
	pass, msg := ValidateRequest(*req)
	if !pass {
		return errors.New(msg)
	}
	User, pass, _ := database.FindUser(req.Email)
	//TODO will deal with passwordHash and not the actual passwords
	if req.Password != User.Password {
		//TODO this static string has to be in constants package
		return errors.New("INCORRECT PASSWORD")
	}
	err, _ = process(req.Email, req.Amount, constants.DEBIT, User)
	return err
}
func Send(reqform string, c *routing.Context) error {
	//User can either tell us the Id or email of the reciever , to send money to
	//will check if the sender is actually a user type and not merchant
	//will check if the reciever is actually a merchant type and not user
	req := &models.SendRequest{}
	err := json.Unmarshal([]byte(reqform), req)
	if err != nil {
		return err
	}
	pass, msg := ValidateSendRequest(*req)
	if !pass {
		return errors.New(msg)
	}
	User, pass, _ := database.FindUser(req.From.Email)
	//TODO will deal with passwordHash and not the actual passwords
	if req.From.Password != User.Password {
		//TODO this static string has to be in constants package
		return errors.New("INCORRECT PASSWORD")
	}
	if User.Type != constants.USER {
		return errors.New(constants.CANNOT_TRANSFER)
	}
	Merchant, pass, _ := database.FindUser(req.To.Email)
	if Merchant.Type != constants.MERCHANT {
		return errors.New(constants.CANNOT_TRANSFER)
	}
	err, Id1 := process(req.From.Email, req.From.Amount, constants.DEBIT, User)
	if err != nil {
		return err
	}
	err, Id2 := process(req.To.Email, req.From.Amount, constants.CREDIT, Merchant)
	primitiveID1, _ := primitive.ObjectIDFromHex(Id1)
	t, _, _ := database.FindTransaction(primitiveID1)
	t.InverseTranactionId = Id2
	database.UpdateTransaction(t)

	primitiveID2, _ := primitive.ObjectIDFromHex(Id2)
	t2, _, _ := database.FindTransaction(primitiveID2)
	t2.InverseTranactionId = Id1
	database.UpdateTransaction(t2)

	return err
}

//USER WILL MAKE A REFUND REQUEST IN RESPONSE OF WHICH
//THE MERCHANT CAN MAKE A RESPONSE TO REFUND
func RefundResponse(reqform string, c *routing.Context) error {

	req := &models.RefundResponse{}
	err := json.Unmarshal([]byte(reqform), req)
	if err != nil {
		return err
	}
	pass, msg := ValidateRefundResponse(*req)
	if !pass {
		return errors.New(msg)
	}

	CancelTransaction(req.TransactionID, true)
	InverseId, _ := primitive.ObjectIDFromHex(req.TransactionID)
	t, _, _ := database.FindTransaction(InverseId)
	CancelTransaction(t.InverseTranactionId, false)

	return nil
}

func History(c *routing.Context) error {

	fmt.Fprintf(c, "returning this string")
	email := string(c.Request.Header.Peek("email"))
	password := string(c.Request.Header.Peek("password"))

	User, found, err := database.FindUser(email)
	if err != nil {
		fmt.Fprintf(c, fmt.Sprintf("%v", err))
		return nil
	}
	if !found {
		fmt.Fprintf(c, "user not found")
		return nil
	}
	if User.Password != password {
		fmt.Fprintf(c, "Incorrect password")
		return nil
	}
	LastTransactions := []models.Transaction{}
	for _, Id := range User.History {
		t, _, _ := database.FindTransaction(Id)
		LastTransactions = append(LastTransactions, t)
	}

	his := fmt.Sprintf("%v", LastTransactions)
	fmt.Fprintf(c, his)
	return nil

}
