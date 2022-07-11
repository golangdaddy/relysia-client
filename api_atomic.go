package relysia

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type OfferRequest struct {
	DataArray []*Offer `json:"dataArray"`
}

type Offer struct {
	Amount  float64 `json:"amount"`
	Type    string  `json:"type"`
	TokenID string  `json:"tokenId"`
	Sn      int     `json:"sn"`
}

type OfferResponse struct {
	Status   string   `json:"status"`
	Msg      string   `json:"msg"`
	Contents []string `json:"contents"`
}

type OfferResponseContent struct {
	MakerPublicKeyHash string  `json:"makerPublicKeyHash"`
	PrevTxid           string  `json:"prevTxid"`
	SerialNumber       float64 `json:"serialNumber"`
	SwapId             string  `json:"swapId"`
	SwapOfferHex       string  `json:"swapOfferHex"`
	TokenId            string  `json:"tokenId"`
	TokenOwnerAddress  string  `json:"tokenOwnerAddress"`
	TokenSatoshis      float64 `json:"tokenSatoshis"`
	WantedSatoshis     float64 `json:"wantedSatoshis"`
}

func (self *Client) Offer(walletID, tokenID, reciveType string, amount float64) (*OfferResponse, error) {
	methodName := "Offer"

	headers := Headers{
		"walletID": walletID,
	}

	b, err := json.Marshal(
		OfferRequest{
			DataArray: []*Offer{
				&Offer{
					Type:    reciveType,
					TokenID: tokenID,
					Amount:  amount,
					Sn:      0,
				},
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	b, err = self.do(
		"POST",
		"v1/offer",
		bytes.NewBuffer(b),
		self.POSTHeaders(headers),
	)
	println(string(b))
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &OfferResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response, nil
}

type SwapRequest struct {
	DataArray []*Swap `json:"dataArray"`
}

type Swap struct {
	SwapHex string `json:"swapHex"`
}

type SwapResponse struct {
	Status string   `json:"status"`
	Msg    string   `json:"msg"`
	TxIds  []string `json:"txIds"`
}

func (self *Client) Swap(walletID, swapHex string) (*SwapResponse, error) {
	methodName := "Swap"

	headers := Headers{
		"walletID": walletID,
	}

	b, err := json.Marshal(
		SwapRequest{
			DataArray: []*Swap{
				&Swap{
					SwapHex: swapHex,
				},
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	b, err = self.do(
		"POST",
		"v1/swap",
		bytes.NewBuffer(b),
		self.POSTHeaders(headers),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &SwapResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response, nil
}
