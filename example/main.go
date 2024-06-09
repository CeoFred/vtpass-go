package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vt "github.com/CeoFred/vtupass_go"

	"github.com/joho/godotenv"
)

var apiKey, publicKey, secretKey string
var service *vt.VTService

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey = os.Getenv("API_KEY")
	publicKey = os.Getenv("PUBLIC_KEY")
	secretKey = os.Getenv("SECRET_KEY")

}

func main() {

	service = vt.NewVTService(apiKey, publicKey, secretKey, vt.EnvironmentSandbox)
	available, err := service.Ping(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("service available ==>", available)

	// Balance()
	// ServiceCategories()
	VerifyMeterNumber()
}

func VerifyMeterNumber() {
	customer, err := service.VerifyMeterNumber(context.Background(), "1111111111111", "prepaid", "enugu-electric")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Name: ", customer.CustomerName)
	fmt.Println("Business Unit: ", customer.BusinessUnit)
	fmt.Println("Account_Number: ", customer.AccountNumber)

}

func ServiceVariations() {
	services, err := service.ServiceVariations(context.Background(), "enugu-electric")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(services); i++ {
		fmt.Printf("Service: %s, Min Amount: %s \n", services[i].Name, services[i].VariationCode)
	}
}

func Balance() {
	walletBalance, err := service.Balance(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("wallet balance: %f\n", walletBalance.Contents.Balance)
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
		fmt.Printf("Service: %s, Min Amount: %s \n", services[i].Name, services[i].MaximumAmount)
	}
}
