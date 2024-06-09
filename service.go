package go_library_starter_kit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	httpclient "github.com/CeoFred/go_library_starter_kit/lib"
)

type Service struct {
	apiKey          string
	client          HttpClient
	authCredentials string
}

const BaseURL = "https://some-baseurl.com"

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	BaseResponse
	Error string `json:"error"`
}

func (e ErrorResponse) Message() string {
	return e.BaseResponse.Message
}

type UserDataResponse struct {
	BaseResponse
	User struct {
		FirstName  string `json:"first_name"`
	} `json:"data"`
}

type Response struct {
	BaseResponse
	Data struct {
		User string `json:"user"`
	} `json:"data"`
}

func NewService(apiKey string) *Service {
	return &Service{
		apiKey:          apiKey,
		client:          httpclient.NewAPIClient(BaseURL, ""),
		authCredentials: fmt.Sprintf("?access_token=%s", apiKey),
	}
}

func (c *Service) GetData(ctx context.Context) (*UserDataResponse, error) {

	resp, err := c.client.Get(ctx, "user"+c.authCredentials)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorResponse.Error)
	}

	var balance UserDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&balance); err != nil {
		return nil, err
	}

	return &balance, nil

}

func (c *Service) PostData(ctx context.Context, payload interface{}) (*Response, error) {

	path := fmt.Sprintf("user%s", c.authCredentials)

	resp, err := c.client.Post(ctx, path, payload)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorResponse.Error)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("error decoding response", err)
		return nil, err
	}

	return &response, nil
}
