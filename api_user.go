package relysia

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserResponse struct {
	Status      string       `json:"status"`
	Msg         string       `json:"msg"`
	UserDetails *UserDetails `json:"userDetails"`
}

type UserDetails struct {
	LocalID           string      `json:"localId"`
	PasswordHash      string      `json:"passwordHash"`
	PasswordUpdatedAt int64       `json:"passwordUpdatedAt"`
	ValidSince        string      `json:"validSince"`
	LastLoginAt       string      `json:"lastLoginAt"`
	CreatedAt         string      `json:"createdAt"`
	LastRefreshAt     time.Time   `json:"lastRefreshAt"`
	PhoneNumber       interface{} `json:"phoneNumber"`
}

func (self *Client) User() (*UserDetails, error) {
	methodName := "User"

	b, err := self.do(
		"GET",
		"v1/user",
		nil,
		self.GETHeaders(),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &UserResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	println(string(b))

	return response.UserDetails, nil
}

func (self *Client) DeleteUser() error {
	methodName := "User"

	_, err := self.do(
		"DELETE",
		"v1/user",
		nil,
		self.GETHeaders(),
	)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return nil
}
