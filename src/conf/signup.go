package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/PankajKumar9/PaymentAPI/src/conf/constants"
	"github.com/PankajKumar9/PaymentAPI/src/database"
	"github.com/PankajKumar9/PaymentAPI/src/models"
	"github.com/PankajKumar9/PaymentAPI/src/utility"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	routing "github.com/qiangxue/fasthttp-routing"
)

const userPwPepper = "abcd"

func Signup(signupform string, c *routing.Context) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	newUser := &models.User{}
	err := json.Unmarshal([]byte(signupform), newUser)
	if err != nil {
		return err
	}

	pass, msg := ValidateUser(*newUser)
	if !pass {
		return errors.New(msg)
	}
	pwBytes := []byte(newUser.Password + userPwPepper)

	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	newUser.PasswordHash = string(hashedBytes)
	newUser.Password = ""
	log.Println(utility.Debug(*newUser))
	//todo check password == confirm password

	//TODO will send password for hashing , wont write the password

	database.InsertUser(*newUser)
	return nil
}

func Credit(reqform string, c *routing.Context) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(utility.Info(reqform))
	req := &models.CreditRequest{}
	err := json.Unmarshal([]byte(reqform), req)
	if err != nil {
		return err
	}
	log.Println(utility.Info(*req))
	pass, msg := ValidateRequest(*req)
	if !pass {
		return errors.New(msg)
	}
	User, pass, err := database.FindUser(req.Email)
	if err != nil {
		log.Println(utility.Info("error in finding this user"))
	}
	if !pass {
		log.Println(utility.Info("This user does not even exists"))
	}
	log.Println(utility.Info(User))
	//TODO will deal with passwordHash and not the actual passwords

	if !ValidatePassword(req.Password, User.PasswordHash) {
		//TODO this static string has to be in constants package

		return errors.New(constants.INCORRECT_PASSWORD)
	}
	err, _ = process(req.Email, req.Amount, constants.CREDIT, User, User.Email, User.Email, constants.CREDIT)
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
	if !ValidatePassword(req.Password, User.PasswordHash) {
		//TODO this static string has to be in constants package
		return errors.New(constants.INCORRECT_PASSWORD)
	}
	err, _ = process(req.Email, req.Amount, constants.DEBIT, User, User.Email, User.Email, constants.DEBIT)
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(utility.Info(*req))
	User, pass, _ := database.FindUser(req.From.Email)
	//TODO will deal with passwordHash and not the actual passwords
	if !ValidatePassword(req.From.Password, User.PasswordHash) {
		//TODO this static string has to be in constants package
		return errors.New(constants.INCORRECT_PASSWORD)
	}
	if User.Type != constants.USER {
		return errors.New(constants.INVALID_USER_TYPE)
	}
	Merchant, pass, _ := database.FindUser(req.To.Email)
	if Merchant.Type != constants.MERCHANT {
		return errors.New(constants.INVALID_MERCHANT_TYPE)
	}
	err, Id1 := process(req.From.Email, req.From.Amount, constants.DEBIT, User, req.From.Email, req.To.Email, constants.PAYMENT)
	if err != nil {
		return err
	}
	err, Id2 := process(req.To.Email, req.From.Amount, constants.CREDIT, Merchant, req.From.Email, req.To.Email, constants.PAYMENT)
	primitiveID1, _ := primitive.ObjectIDFromHex(Id1)
	t, _, _ := database.FindTransaction(primitiveID1)
	t.InverseTranactionId = Id2
	database.UpdateTransaction(*t)

	primitiveID2, _ := primitive.ObjectIDFromHex(Id2)
	t2, _, _ := database.FindTransaction(primitiveID2)
	t2.InverseTranactionId = Id1
	database.UpdateTransaction(*t2)

	return err
}

//USER WILL MAKE A REFUND REQUEST IN RESPONSE OF WHICH
//THE MERCHANT CAN MAKE A RESPONSE TO REFUND
func RefundResponse(reqform string, c *routing.Context) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	req := &models.RefundResponse{}
	err := json.Unmarshal([]byte(reqform), req)
	if err != nil {
		return err
	}
	pass, msg := ValidateRefundResponse(*req)
	if !pass {
		return errors.New(msg)
	}
	RefundId, _ := primitive.ObjectIDFromHex(req.TransactionID)
	actualTransaction, pass, err := database.FindTransaction(RefundId)
	if !pass {
		log.Println(utility.Info("didnt pass"))
	}
	if err != nil {
		log.Println(utility.Info(err))
	}
	log.Println(utility.Info(req.TransactionID))
	log.Println(utility.Info(RefundId))
	//TODO PUSH THIS LOGIC IN CANCEL FUNCTION SO
	//DONT NEED TO  REDUNTIVELY FIND THE ACTUAL TRANSACTION
	merchant, _, _ := database.FindUser(req.Email)
	if !ValidatePassword(req.Password, merchant.PasswordHash) {
		return errors.New(constants.INCORRECT_PASSWORD)
	}
	if actualTransaction.Info.To != req.Email {
		log.Println(utility.Info(actualTransaction.Info.To))
		log.Println(utility.Info(req.Email))
		return errors.New("THIS MERCHANT DID NOT RECIEVE THE SAME TRANSACTION")
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
	if !ValidatePassword(password, User.PasswordHash) {
		return errors.New(constants.INCORRECT_PASSWORD)
	}

	LastTransactions := []models.Transaction{}
	for _, Id := range User.History {
		t, _, _ := database.FindTransaction(Id)
		LastTransactions = append(LastTransactions, *t)
	}

	his := fmt.Sprintf("%v", LastTransactions)
	fmt.Fprintf(c, his)
	return nil

}
