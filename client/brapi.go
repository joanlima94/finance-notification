package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"finance-notification/model"
)

const baseURL = "https://brapi.dev/api"

type BrapiClient struct {
	token      string
	httpClient *http.Client
}

func NewBrapiClient(token string) *BrapiClient {
	return &BrapiClient{
		token:      token,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *BrapiClient) doRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.token)
	return c.httpClient.Do(req)
}

func (c *BrapiClient) FetchAllTickers() ([]model.Stock, error) {
	resp, err := c.doRequest(fmt.Sprintf("%s/quote/list", baseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status inesperado: %s", resp.Status)
	}

	var listResp model.StockListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Stocks, nil
}

func (c *BrapiClient) FetchQuote(ticker string) (*model.Quote, error) {
	url := fmt.Sprintf("%s/quote/%s?interval=1d&range=1d", baseURL, ticker)

	resp, err := c.doRequest(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status inesperado: %s", resp.Status)
	}

	var brapiResp model.BrapiResponse
	if err := json.NewDecoder(resp.Body).Decode(&brapiResp); err != nil {
		return nil, err
	}

	if len(brapiResp.Results) == 0 {
		return nil, nil
	}

	return &brapiResp.Results[0], nil
}