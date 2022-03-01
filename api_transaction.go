package relysia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

func (self *Client) Send(walletID, address, currencyType, notes string, amount float64) error {
	methodName := "Send"

	headers := Headers{
		"walletID": walletID,
	}

	b, _ := json.Marshal(
		map[string]string{
			"address": address,
			"notes":   notes,
			"type":    currencyType,
			"amount":  strconv.FormatFloat(amount, 'f', -1, 64),
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
