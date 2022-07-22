# PaymentAPI

- # User can send the money into merchant
there is a send api, in the body ,you need to send a json something like this :
{
    "from ": {
        "email ": "",
        "password": "",
        "amount": 100
    },
    "to": {
        "email": "",
        "Id": ""
    }
}

- # Merchant can send the money back to user (Refund)
there is a refund api in the body ,you need to send a json something like this :
{
    "email ": "",
    "password": "",
    "transactionId": ""
}
- # Merchant can withdraw money from her account
there is a debit apit for any user can withdraw money from their own account account ,you need to send a json something like this :
{
    "email ": "",
    "password": "",
    "amount": ""
}
- # Merchant can check the transactions related to her account
any user can see thier transaction history with a GET request 
in the headers they need to provide their email and password 

others:
- # the Payment App should keep track of all the merchant and users transactions
any transaction of anykind, has extensive record in the database