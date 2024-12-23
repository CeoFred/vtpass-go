package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vt "github.com/CeoFred/vtpass-go"

	"github.com/joho/godotenv"
)

var apiKey, publicKey, secretKey, env string
var service *vt.VTService

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey = os.Getenv("API_KEY")
	publicKey = os.Getenv("PUBLIC_KEY")
	secretKey = os.Getenv("SECRET_KEY")
	env = os.Getenv("ENVIRONMENT")

}

func main() {

	service = vt.NewVTService(apiKey, publicKey, secretKey, vt.Environment(env))
	available, err := service.Ping(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("service available ==>", available)

	// ServiceByIdentifier()
	// ServiceByIdentifier()
	// PayElectricityPrepaid()
	QueryTransaction()
}

func QueryTransaction() {
	// id := service.GenerateRequestID()

	// re, err := service.PurchaseElectricity(context.Background(), vt.ElectricityPurchase{
	// 	RequestID:     id,
	// 	ServiceID:     "enugu-electric",
	// 	BillersCode:   "1010101010101",
	// 	VariationCode: "postpaid",
	// 	Amount:        1000,
	// 	Phone:         "08160583193",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }

	txn, err := service.QueryTransaction(context.Background(), "202412231011dc5795c1b4f040f78bf2333e84f15e8b")
	if err != nil {
		panic(err)
	}

	fmt.Println(txn)
}

func PayElectricityPostpaid() {
	id := service.GenerateRequestID()

	response, err := service.PurchaseElectricity(context.Background(), vt.ElectricityPurchase{
		RequestID:     id,
		ServiceID:     "enugu-electric",
		BillersCode:   "1010101010101",
		VariationCode: "postpaid",
		Amount:        70.23,
		Phone:         "08160583193",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("postpaid purchase \n", response.Content)

}

func PayElectricityPrepaid() {
	id := service.GenerateRequestID()

	response, err := service.PurchaseElectricity(context.Background(), vt.ElectricityPurchase{
		RequestID:     id,
		ServiceID:     "portharcourt-electric",
		BillersCode:   "1111111111111",
		VariationCode: "prepaid",
		Amount:        1000,
		Phone:         "8160583193",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("prepaid purchase code \n", response.PurchasedCode)
	fmt.Println(response.Content.Transactions.UnitPrice)
	fmt.Println("code => ", response.Code)
	fmt.Println("ID => ", id)



}

func GenerateRequestID() {
	requestID := service.GenerateRequestID()

	fmt.Println("Request ID: ", requestID)
}

func VerifyMeterNumber() {
	customer, err := service.VerifyMeterNumber(context.Background(), "0137200395333", "prepaid", "portharcourt-electric")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Customer: ", customer)
	fmt.Println("Business Unit: ", customer.BusinessUnit)
	fmt.Println("Account_Number: ", customer.AccountNumber)

}

func ServiceVariations() {
	services, err := service.ServiceVariations(context.Background(), "enugu-electric")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(services); i++ {
		fmt.Printf("Code: %s, Min Amount: %s , Variation amount: %s \n", services[i].VariationCode, services[i].FixedPrice, services[i].VariationAmount)
	}
}

func Balance() {
	walletBalance, err := service.Balance(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	if walletBalance != nil {
		fmt.Printf("wallet balance: %s\n", walletBalance.Contents.Balance)
	}
}

func ServiceCategories() {
	services, err := service.ServiceCategories(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(services); i++ {
		fmt.Println("Service: ", services[i].Name)
	}
}

func ServiceByIdentifier() {
	services, err := service.ServiceByIdentifier(context.Background(), vt.IdentifierElectricityBill)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(services); i++ {
		fmt.Printf("Service: %s, Min Amount: %s \n", services[i].Name, services[i].MinimumAmount)
	}
}
