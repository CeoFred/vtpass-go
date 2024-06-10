package vtupass_go

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	httpclient "github.com/CeoFred/vtpass_go/lib"

	"github.com/google/uuid"
)

type VTService struct {
	apiKey          string
	publicKey       string
	secretKey       string
	client          HttpClient
	authCredentials map[string]string
	Enviroment      Environment
}

type BaseResponse struct {
	Code string `json:"code"`
}

type ErrorResponse struct {
	BaseResponse
}

func (e ErrorResponse) Error() string {

	var message string

	switch e.Code {
	case BILLER_CONFIRMED:
		message = "BILLER CONFIRMED"
	case INVALID_ARGUMENTS:
		message = "INVALID ARGUMENTS"
	case PRODUCT_DOES_NOT_EXIST:
		message = "PRODUCT_DOES_NOT_EXIST"
	case BILLER_NOT_REACHABLE_AT_THIS_POINT:
		message = "BILLER NOT REACHABLE AT THIS POINT"
	}
	return message
}

type WalletBalance struct {
	BaseResponse
	Contents struct {
		Balance float64 `json:"balance"`
	} `json:"contents"`
}

type ServiceCategoryResponse struct {
	BaseResponse
	Content             []ServiceCategory `json:"content"`
	ResponseDescription string            `json:"response_description"`
}

type ServiceResponse struct {
	BaseResponse
	Content             []Service `json:"content"`
	ResponseDescription string    `json:"response_description"`
}

type VariationResponse struct {
	BaseResponse
	Content struct {
		ServiceName string      `json:"ServiceName"`
		Variations  []Variation `json:"varations"`
	} `json:"content"`
}

type CustomerInfoResponse struct {
	BaseResponse
	Content CustomerInfo `json:"content"`
}

func NewVTService(apiKey, publicKey, secretKey string, environment Environment) *VTService {
	var baseUrl string

	switch environment {
	case EnvironmentSandbox:
		baseUrl = SandboxBaseURL
	case EnvironmentLive:
		baseUrl = LiveEnviromentURL
	default:
		baseUrl = SandboxBaseURL
	}

	return &VTService{
		apiKey:     apiKey,
		client:     httpclient.NewAPIClient(baseUrl, apiKey),
		Enviroment: environment,
		publicKey:  publicKey,
		secretKey:  secretKey,
		authCredentials: map[string]string{
			"api-key":    apiKey,
			"public-key": publicKey,
			"secret-key": secretKey,
		},
	}
}

type Details struct {
	AppliedToArrears  float64 `json:"appliedToArrears"`
	ArrearsBalance    float64 `json:"arrearsBalance"`
	Wallet            float64 `json:"wallet"`
	ExchangeReference string  `json:"exchangeReference"`
	VAT               float64 `json:"vat"`
	InvoiceNumber     string  `json:"invoiceNumber"`
	AppliedToWallet   float64 `json:"appliedToWallet"`
	Units             float64 `json:"units"`
	ResponseMessage   string  `json:"responseMessage"`
	Status            string  `json:"status"`
	ResponseCode      int     `json:"responseCode"`
	Token             string  `json:"token"`
}

type TransactionResponse struct {
	Code    string `json:"code"`
	Content struct {
		Details           Details `json:"details"`
		TransactionNumber string  `json:"transactionNumber"`
	} `json:"content"`
}

// QUERY TRANSACTION STATUS
func (s *VTService) QueryTransaction(ctx context.Context, request_id string) (*TransactionResponse, error) {
	url := "requery"

	payload := map[string]interface{}{
		"request_id": request_id,
	}

	resp, err := s.client.Post(ctx, url, payload, s.authCredentials)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}
		return nil, errorResponse
	}

	var resonse TransactionResponse
	if err := json.NewDecoder(resp.Body).Decode(&resonse); err != nil {
		return nil, err
	}

	if resonse.Code == "011" {
		return nil, fmt.Errorf("service not valid or invalid argumments")
	}
	if resonse.Code == "012" {
		return nil, fmt.Errorf("prodduct does not exist")
	}

	return &resonse, nil

}

// PURCHASE PRODUCT (Payment)
// https://www.vtpass.com/documentation/eedc-enugu-electric-api/
func (s *VTService) PurchaseElectricity(ctx context.Context, payload ElectricityPurchase) (*PayResponse, error) {

	url := "pay"
	resp, err := s.client.Post(ctx, url, payload, s.authCredentials)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, errorResponse
	}

	var resonse PayResponse
	if err := json.NewDecoder(resp.Body).Decode(&resonse); err != nil {
		return nil, err
	}

	if resonse.Code == "011" {
		return nil, fmt.Errorf("service not valid or invalid argumments")
	}
	if resonse.Code == "012" {
		return nil, fmt.Errorf("prodduct does not exist")
	}
	return &resonse, nil

}

// REQUEST ID
// https://www.vtpass.com/documentation/how-to-generate-request-id/
func (s *VTService) GenerateRequestID() string {
	location, err := time.LoadLocation("Africa/Lagos")
	if err != nil {
		// Fallback to local time if the location is not found
		location = time.Now().Location()
	}

	// Get the current time in Africa/Lagos
	currentTime := time.Now().In(location)

	// Format the time to YYYYMMDDHHII
	timeFormatted := currentTime.Format("200601021504")

	// Generate a UUID
	uuidPart := uuid.New().String()

	// Concatenate the formatted time and UUID
	requestID := timeFormatted + "-" + uuidPart

	requestID = strings.ReplaceAll(requestID, "-", "")
	return requestID
}

// VERIFY METER NUMBER
// https://www.vtpass.com/documentation/eedc-enugu-electric-api/
func (s *VTService) VerifyMeterNumber(ctx context.Context, meter_number, meter_type, service_id string) (*CustomerInfo, error) {
	url := "merchant-verify"

	requestData := map[string]interface{}{
		"billersCode": meter_number,
		"serviceID":   service_id,
		"type":        meter_type,
	}

	resp, err := s.client.Post(ctx, url, requestData, s.authCredentials)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, errorResponse
	}

	var resonse CustomerInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&resonse); err != nil {
		return nil, err
	}

	if resonse.Code == "011" {
		return nil, fmt.Errorf("service not valid or invalid argumments")
	}
	if resonse.Code == "012" {
		return nil, fmt.Errorf("prodduct does not exist")
	}

	return &resonse.Content, nil

}

// GET VARIATION CODES
// https://www.vtpass.com/documentation/variation-codes/
func (s *VTService) ServiceVariations(ctx context.Context, id string) ([]Variation, error) {
	url := fmt.Sprintf("service-variations?serviceID=%s", id)

	resp, err := s.client.Get(ctx, url, s.authCredentials)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, errorResponse
	}
	var resonse VariationResponse
	if err := json.NewDecoder(resp.Body).Decode(&resonse); err != nil {
		return nil, err
	}

	if resonse.Code == "011" {
		return nil, fmt.Errorf("service not valid or invalid argumments")
	}

	return resonse.Content.Variations, nil
}

// GET SERVICE ID
// https://www.vtpass.com/documentation/service-ids/
func (s *VTService) ServiceByIdentifier(ctx context.Context, id string) ([]Service, error) {
	url := fmt.Sprintf("services?identifier=%s", id)

	resp, err := s.client.Get(ctx, url, s.authCredentials)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, errorResponse
	}
	var resonse ServiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&resonse); err != nil {
		return nil, err
	}
	if resonse.Code == "011" {
		return nil, fmt.Errorf("service not valid or invalid argumments")
	}

	return resonse.Content, nil
}

func (s *VTService) ServiceCategories(ctx context.Context) ([]ServiceCategory, error) {
	resp, err := s.client.Get(ctx, "service-categories", s.authCredentials)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, errorResponse
	}
	var resonse ServiceCategoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&resonse); err != nil {
		return nil, err
	}
	if resonse.Code == "011" {
		return nil, fmt.Errorf("service not valid or invalid argumments")
	}

	return resonse.Content, nil
}

// Test authentication
func (s *VTService) Ping(ctx context.Context) (bool, error) {
	resp, err := s.client.Get(ctx, "balance", s.authCredentials)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return false, err
		}

		return false, errorResponse
	}
	return true, nil
}

func (s *VTService) Balance(ctx context.Context) (*WalletBalance, error) {

	resp, err := s.client.Get(ctx, "balance", s.authCredentials)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, errorResponse
	}

	var resonse WalletBalance
	if err := json.NewDecoder(resp.Body).Decode(&resonse); err != nil {
		return nil, err
	}
	if resonse.Code == "011" {
		return nil, fmt.Errorf("service not valid or invalid argumments")
	}
	return &resonse, nil

}

// func (c *Service) PostData(ctx context.Context, payload interface{}) (*Response, error) {

// 	path := fmt.Sprintf("user%s", c.authCredentials)

// 	resp, err := c.client.Post(ctx, path, payload)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		var errorResponse ErrorResponse
// 		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
// 			return nil, err
// 		}

// 		return nil, errorResponse
// 	}

// 	var response Response
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("error decoding response", err)
// 		return nil, err
// 	}

// 	return &response, nil
// }
