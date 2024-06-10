# VTPass Go Library Documentation

## Overview

The VTPass Go library provides an interface to interact with the VTPass API for various services such as querying transaction status, purchasing products (electricity), verifying meter numbers, and fetching service variations and categories.

## Installation

To install the VTPass Go library, use the following command:

```sh
go get github.com/CeoFred/vtpass-go
```

## Usage

First, initialize the service by providing your API credentials and environment:

```go
import (
    "context"
    "fmt"
    "log"
    "os"

    vt "github.com/CeoFred/vtpass-go"
    "github.com/joho/godotenv"
)

var apiKey, publicKey, secretKey string
var service *vt.VTService

func init() {
    if err := godotenv.Load(".env"); err != nil {
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
    fmt.Println("service available:", available)

    // Example usage
    CheckBalance()
    PurchaseElectricityPrepaid()
    VerifyMeterNumber()
}

func CheckBalance() {
    walletBalance, err := service.Balance(context.Background())
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("wallet balance: %f\n", walletBalance.Contents.Balance)
}

func PurchaseElectricityPrepaid() {
    id := service.GenerateRequestID()
    response, err := service.PurchaseElectricity(context.Background(), vt.ElectricityPurchase{
        RequestID:    id,
        ServiceID:    "enugu-electric",
        BillersCode:  "1111111111111",
        VariationCode: "prepaid",
        Amount:       1000,
        Phone:        "8160583193",
    })
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("prepaid purchase code:", response.PurchasedCode)
}

func VerifyMeterNumber() {
    customer, err := service.VerifyMeterNumber(context.Background(), "1111111111111", "prepaid", "enugu-electric")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Name:", customer.CustomerName)
    fmt.Println("Business Unit:", customer.BusinessUnit)
    fmt.Println("Account Number:", customer.AccountNumber)
}
```

## Service Methods

### `NewVTService(apiKey, publicKey, secretKey string, environment Environment) *VTService`
Creates a new instance of the VTService with the provided API credentials and environment (sandbox or live).

### `Ping(ctx context.Context) (bool, error)`
Checks the service availability.

**Example Usage:**

```go
available, err := service.Ping(context.Background())
if err != nil {
    fmt.Println(err)
}
fmt.Println("service available:", available)
```

### `Balance(ctx context.Context) (*WalletBalance, error)`
Fetches the wallet balance.

**Example Usage:**

```go
walletBalance, err := service.Balance(context.Background())
if err != nil {
    fmt.Println(err)
}
fmt.Printf("wallet balance: %f\n", walletBalance.Contents.Balance)
```

### `GenerateRequestID() string`
Generates a unique request ID based on the current time and a UUID.

**Example Usage:**

```go
requestID := service.GenerateRequestID()
fmt.Println("Request ID:", requestID)
```

### `QueryTransaction(ctx context.Context, request_id string) (*TransactionResponse, error)`
Queries the status of a transaction using the request ID.

**Example Usage:**

```go
txn, err := service.QueryTransaction(context.Background(), "request_id_here")
if err != nil {
    fmt.Println(err)
}
fmt.Println("Transaction:", txn)
```

### `PurchaseElectricity(ctx context.Context, payload ElectricityPurchase) (*PayResponse, error)`
Purchases electricity by providing the necessary details.

**Example Usage:**

```go
id := service.GenerateRequestID()
response, err := service.PurchaseElectricity(context.Background(), vt.ElectricityPurchase{
    RequestID:    id,
    ServiceID:    "enugu-electric",
    BillersCode:  "1111111111111",
    VariationCode: "prepaid",
    Amount:       1000,
    Phone:        "8160583193",
})
if err != nil {
    fmt.Println(err)
}
fmt.Println("prepaid purchase code:", response.PurchasedCode)
```

### `VerifyMeterNumber(ctx context.Context, meter_number, meter_type, service_id string) (*CustomerInfo, error)`
Verifies a meter number.

**Example Usage:**

```go
customer, err := service.VerifyMeterNumber(context.Background(), "1111111111111", "prepaid", "enugu-electric")
if err != nil {
    fmt.Println(err)
}
fmt.Println("Name:", customer.CustomerName)
fmt.Println("Business Unit:", customer.BusinessUnit)
fmt.Println("Account Number:", customer.AccountNumber)
```

### `ServiceVariations(ctx context.Context, id string) ([]Variation, error)`
Fetches the service variations for a given service ID.

**Example Usage:**

```go
services, err := service.ServiceVariations(context.Background(), "enugu-electric")
if err != nil {
    fmt.Println(err)
}
for _, service := range services {
    fmt.Printf("Service: %s, Min Amount: %s\n", service.Name, service.VariationCode)
}
```

### `ServiceByIdentifier(ctx context.Context, id string) ([]Service, error)`
Fetches services by their identifier.

**Example Usage:**

```go
services, err := service.ServiceByIdentifier(context.Background(), vt.IdentifierElectricityBill)
if err != nil {
    fmt.Println(err)
}
for _, service := range services {
    fmt.Printf("Service: %s, Min Amount: %s\n", service.Name, service.MaximumAmount)
}
```

### `ServiceCategories(ctx context.Context) ([]ServiceCategory, error)`
Fetches all service categories.

**Example Usage:**

```go
services, err := service.ServiceCategories(context.Background())
if err != nil {
    fmt.Println(err)
}
for _, service := range services {
    fmt.Println("Service:", service.Name)
}
```

## Error Handling

All service methods return an error as the second return value. Check this error to handle any issues that arise during the API call.

**Example Usage:**

```go
available, err := service.Ping(context.Background())
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Service available:", available)
}
```