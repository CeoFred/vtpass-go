package vtupass_go


type Environment string

const (
	EnvironmentSandbox Environment = "sandbox"
	EnvironmentLive Environment = "live"
)

type ServiceCategory struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type Service struct {
	ServiceID       string `json:"serviceID"`
	Name            string `json:"name"`
	MinimumAmount   string `json:"minimum_amount"`
	MaximumAmount   string    `json:"maximum_amount"`
	ConvenienceFee  string `json:"convenience_fee"`
	ProductType     string `json:"product_type"`
	Image           string `json:"image"`
}

type Variation struct {
	VariationCode   string  `json:"variation_code"`
	Name            string  `json:"name"`
	VariationAmount string `json:"variation_amount"`
	FixedPrice      string  `json:"fixedPrice"`
}

type CustomerInfo struct {
	CustomerName     string `json:"Customer_Name"`
	AccountNumber    string `json:"Account_Number"`
	MeterNumber      string `json:"Meter_Number"`
	BusinessUnit     string `json:"Business_Unit"`
	Address          string `json:"Address"`
	CustomerArrears  string `json:"Customer_Arrears"`
}