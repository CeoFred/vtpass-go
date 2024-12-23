package vtupass_go

type Environment string

const (
	EnvironmentSandbox Environment = "sandbox"
	EnvironmentLive    Environment = "live"
)

type ServiceCategory struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type Service struct {
	ServiceID      string `json:"serviceID"`
	Name           string `json:"name"`
	MinimumAmount  string `json:"minimium_amount"`
	MaximumAmount  uint   `json:"maximum_amount"`
	ConvenienceFee string `json:"convinience_fee"`
	ProductType    string `json:"product_type"`
	Image          string `json:"image"`
}

type Variation struct {
	VariationCode   string `json:"variation_code"`
	Name            string `json:"name"`
	VariationAmount string `json:"variation_amount"`
	FixedPrice      string `json:"fixedPrice"`
}

type CustomerInfo struct {
	CustomerName     string `json:"Customer_Name"`
	CustomerNumber   string `json:"Customer_Number"`
	AccountNumber    string `json:"Account_Number"`
	MeterNumber      string `json:"MeterNumber"`
	BusinessUnit     string `json:"Business_Unit"`
	Address          string `json:"Address"`
	CustomerArrears  string `json:"Customer_Arrears"`
	District         string `json:"District"`
	LastPurchaseDays string `json:"Last_Purchase_Days"`
	KCT1             string `json:"KCT1"`
	KCT2             string `json:"KCT2"`
	// MAXPurchaseAmount string `json:"Max_Purchase_Amount"`
	// MinPurchaseAmount string `json:"Min_Purchase_Amount"`
	CustomerPhone string `json:"Customer_Phone"`
}

type ElectricityPurchase struct {
	RequestID     string  `json:"request_id"`
	ServiceID     string  `json:"serviceID"`
	BillersCode   string  `json:"billersCode"`
	VariationCode string  `json:"variation_code"`
	Amount        float64 `json:"amount"`
	Phone         string  `json:"phone"`
}

type Data struct {
	Code                string  `json:"code"`
	Content             Content `json:"content"`
	ResponseDescription string  `json:"response_description"`
	Amount              float64 `json:"amount"`
	TransactionDate     *string `json:"transaction_date"`
	RequestID           string  `json:"requestId"`
	PurchasedCode       string  `json:"purchased_code"`
}

type TransactionUpdate struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}
type Transaction struct {
	Amount              interface{}     `json:"amount"`
	ConvenienceFee      interface{}     `json:"convenience_fee"`
	Status              string      `json:"status"`
	Name                *string     `json:"name"`
	Phone               string      `json:"phone"`
	Email               string      `json:"email"`
	Type                string      `json:"type"`
	CreatedAt           string      `json:"created_at"`
	Discount            *string     `json:"discount"`
	GiftcardID          *string     `json:"giftcard_id"`
	TotalAmount         float64     `json:"total_amount"`
	Commission          float64     `json:"commission"`
	Channel             string      `json:"channel"`
	Platform            string      `json:"platform"`
	ServiceVerification *string     `json:"service_verification"`
	Quantity            float64     `json:"quantity"`
	UnitPrice           interface{} `json:"unit_price"`
	UniqueElement       string      `json:"unique_element"`
	ProductName         string      `json:"product_name"`
	TransactionID       string      `json:"transactionId"`
	WalletCreditID      string      `json:"wallet_credit_id"`
}

type Content struct {
	Transactions Transaction `json:"transactions"`
}

type TransactionDate struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

type PayResponse struct {
	Code                  string          `json:"code"`
	Content               Content         `json:"content"`
	ResponseDescription   string          `json:"response_description"`
	RequestID             string          `json:"requestId"`
	Amount                string          `json:"amount"`
	TransactionDate       TransactionDate `json:"transaction_date"`
	PurchasedCode         string          `json:"purchased_code"`
	MainToken             string          `json:"mainToken"`
	MainTokenDescription  string          `json:"mainTokenDescription"`
	MainTokenUnits        float64         `json:"mainTokenUnits"`
	MainTokenTax          float64         `json:"mainTokenTax"`
	MainsTokenAmount      float64         `json:"mainsTokenAmount"`
	BonusToken            string          `json:"bonusToken"`
	BonusTokenDescription string          `json:"bonusTokenDescription"`
	BonusTokenUnits       int             `json:"bonusTokenUnits"`
	BonusTokenTax         *float64        `json:"bonusTokenTax"`
	BonusTokenAmount      *float64        `json:"bonusTokenAmount"`
	TariffIndex           string          `json:"tariffIndex"`
	DebtDescription       string          `json:"debtDescription"`
	ExchangeReference     string          `json:"exchangeReference"`
	UtilityName           string          `json:"utilityName"`
}
