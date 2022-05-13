package relysia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type TokenResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Token  string `json:"token"`
	UserID string `json:"userId"`
}

func (self *Client) SignUp(email, pass string) (string, error) {
	methodName := "SignUp"

	b, _ := json.Marshal(
		map[string]string{
			"email":    email,
			"password": pass,
		},
	)
	if self.insecureMode {
		log.Println("SIGNUP", string(b))
	}
	b, err := self.do(
		"POST",
		"v1/signUp",
		bytes.NewBuffer(b),
		self.POSTHeaders(),
	)
	if err != nil {
		return "", fmt.Errorf("%s: %s", methodName, err.Error())
	}

	tokenResponse := &TokenResponse{}
	if err := json.Unmarshal(b, tokenResponse); err != nil {
		return "", fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return tokenResponse.Token, nil
}

func (self *Client) Auth(email, pass string) (string, error) {
	methodName := "Auth"

	b, _ := json.Marshal(
		map[string]string{
			"email":    email,
			"password": pass,
		},
	)
	if self.insecureMode {
		log.Println("LOGIN", string(b))
	}
	b, err := self.do(
		"POST",
		"v1/auth",
		bytes.NewBuffer(b),
		self.POSTHeaders(),
	)
	if err != nil {
		return "", fmt.Errorf("%s: %s", methodName, err.Error())
	}

	tokenResponse := &TokenResponse{}
	if err := json.Unmarshal(b, tokenResponse); err != nil {
		return "", fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return tokenResponse.Token, nil
}
