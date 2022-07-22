package conf

import (
	"github.com/PankajKumar9/PaymentAPI/src/database"
	"github.com/PankajKumar9/PaymentAPI/src/models"
	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(user models.User) (bool, string) {

	if len(user.Email) == 0 {
		return false, "NO EMAIL SENT"
	}
	_, pass, _ := database.FindUser(user.Email)
	if pass {
		return false, "User already exists"
	}

	//TODO : write logic to check mandetory fields
	return true, ""
}
func ValidateRequest(req models.CreditRequest) (bool, string) {
	//TODO check if all the fields are present and are in valid format
	return true, ""

}
func ValidateSendRequest(req models.SendRequest) (bool, string) {
	//TODO check if all the fields are present and are in valid format
	return true, ""

}
func ValidateRefundResponse(req models.RefundResponse) (bool, string) {
	//TODO check if all the fields are present and are in valid format
	return true, ""

}

func ValidatePassword(inputPassword string, userHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userHash), []byte(inputPassword+userPwPepper))
	if err != nil {
		return false
	}
	return true

}
