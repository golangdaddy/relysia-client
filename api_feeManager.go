package relysia

import (
	"encoding/json"
	"fmt"
)

func (self *Client) InitBeta() error {
	methodName := "InitBeta"

	_, err := self.do(
		"GET",
		"v1/initBeta",
		nil,
		self.GETHeaders(),
	)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return nil
}

func (self *Client) FeeMetricsBeta() error {
	methodName := "FeeMetricsBeta"

	b, err := self.do(
		"GET",
		"v1/feeMetricsBeta",
		nil,
		self.GETHeaders(),
	)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &FeeMetricsBetaResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}

	println(string(b))

	return nil
}

type FeeAddressBetaResponse struct {
	Status    string   `json:"status"`
	Msg       string   `json:"msg"`
	Addresses []string `json:"addresses"`
}

func (self *Client) FeeAddressBeta() ([]string, error) {
	methodName := "FeeAddressBeta"

	b, err := self.do(
		"GET",
		"v1/feeAddressBeta",
		nil,
		self.GETHeaders(),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &FeeAddressBetaResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	println(string(b))

	return response.Addresses, nil
}

type FeeMetricsBetaResponse struct {
	StatusCode int `json:"statusCode"`
	Data       struct {
		Msg          string `json:"msg"`
		Status       string `json:"status"`
		TotalBalance struct {
			Balance int `json:"balance"`
		} `json:"totalBalance"`
		FeeUtxos map[int][]*UTXO `json:"feeUtxos"`
	} `json:"data"`
}

type UTXO struct {
	Height int    `json:"height"`
	TxPos  int    `json:"tx_pos"`
	TxHash string `json:"tx_hash"`
	Value  int    `json:"value"`
	Script string `json:"script"`
}
