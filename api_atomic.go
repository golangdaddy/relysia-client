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

func (self *Client) Offer(walletID, tokenID, reciveType string, amount float64) (map[string]interface{}, error) {
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
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := map[string]interface{}{}
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response, nil
}
