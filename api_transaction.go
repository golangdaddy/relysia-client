package relysia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func (self *Client) Send(walletID, address string, amount float64) error {
	methodName := "Send"

	headers := Headers{
		"walletID": walletID,
	}

	b, _ := json.Marshal(
		map[string]interface{}{
			"dataArray": []map[string]interface{}{
				map[string]interface{}{
					"to":     address,
					"amount": amount,
				},
			},
		},
	)
	println("SEND", string(b))

	b, err := self.do(
		"POST",
		"v1/send",
		bytes.NewBuffer(b),
		self.POSTHeaders(headers),
	)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}

	var response interface{} = new(interface{})
	if err := json.Unmarshal(b, response); err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return nil
}

// {"status":"success","msg":"operation completed successfully","txIds":["6ff82c5e1d1e7b4e9e1da40c45363e8ad241d81d712bb47d29c2241837732263"]}
type SendTokenResponse struct {
	Status string   `json:"status"`
	Msg    string   `json:"msg"`
	TxIds  []string `json:"txIds"`
}

func (self *Client) SendToken(walletID, address, tokenID string, amount float64) (*SendTokenResponse, error) {
	methodName := "SendToken"

	headers := Headers{
		"walletID": walletID,
	}

	b, _ := json.Marshal(
		map[string]interface{}{
			"dataArray": []map[string]interface{}{
				map[string]interface{}{
					"to":      address,
					"amount":  amount,
					"tokenId": tokenID,
				},
			},
		},
	)
	println("SEND TOKEN", string(b))

	b, err := self.do(
		"POST",
		"v1/send",
		bytes.NewBuffer(b),
		self.POSTHeaders(headers),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	log.Println(string(b))

	response := &SendTokenResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response, nil
}
