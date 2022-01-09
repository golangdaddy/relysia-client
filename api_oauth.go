package relysia

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (self *Client) OauthToken(clientKey, clientSecret, code string) error {
	methodName := "OauthToken"

	payload := map[string]string{
		"clientKey":    clientKey,
		"clientSecret": clientSecret,
		"code":         code,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}
	println(string(b))
	_, err = self.do(
		"POST",
		"v1/oauth/token",
		bytes.NewBuffer(b),
		self.POSTHeaders(),
	)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}
	return nil
}
