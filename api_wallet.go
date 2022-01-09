package relysia

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (self *Client) CreateWallet(walletTitle string) error {
	methodName := "CreateWallet"

	headers := Headers{
		"walletTitle": walletTitle,
	}

	_, err := self.do(
		"GET",
		"v1/createWallet",
		nil,
		self.GETHeaders(headers),
	)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return nil
}

type WalletsResponse struct {
	Status  string        `json:"status"`
	Msg     string        `json:"msg"`
	Wallets []*WalletInfo `json:"wallets"`
}

type WalletInfo struct {
	WalletID    string `json:"walletID"`
	WalletTitle string `json:"walletTitle"`
	WalletLogo  string `json:"walletLogo,omitempty"`
}

func (self *Client) Wallets() ([]*WalletInfo, error) {
	methodName := "Wallets"

	b, err := self.do(
		"GET",
		"v1/wallets",
		nil,
		self.GETHeaders(),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &WalletsResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response.Wallets, nil
}

type MnemonicResponse struct {
	Status   string `json:"status"`
	Msg      string `json:"msg"`
	Mnemonic string `json:"mnemonic"`
}

func (self *Client) Mnemonic() (string, error) {
	methodName := "Mnemonic"

	b, err := self.do(
		"GET",
		"v1/mnemonic",
		nil,
		self.GETHeaders(),
	)
	if err != nil {
		return "", fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &MnemonicResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return "", fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response.Mnemonic, nil
}

type AddressResponse struct {
	Status  string `json:"status"`
	Msg     string `json:"msg"`
	Address string `json:"address"`
	Paymail string `json:"paymail"`
}

func (self *Client) Address(walletID string) (string, string, error) {
	methodName := "Address"

	headers := Headers{
		"walletID": walletID,
	}

	b, err := self.do(
		"GET",
		"v1/address",
		nil,
		self.GETHeaders(headers),
	)
	if err != nil {
		return "", "", fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &AddressResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return "", "", fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response.Address, response.Paymail, nil
}

type CurrencyResponse struct {
	Status   string  `json:"status"`
	Msg      string  `json:"msg"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}

func (self *Client) CurrencyConversion(satoshis int, currency string) (float64, error) {
	methodName := "CurrencyConversion"

	headers := Headers{
		"satoshis": strconv.Itoa(satoshis),
		"currency": currency,
	}

	b, err := self.do(
		"GET",
		"v1/currencyConversion",
		nil,
		self.GETHeaders(headers),
	)
	if err != nil {
		return 0, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &CurrencyResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return 0, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response.Balance, nil
}

func (self *Client) DeleteWallet(id string) error {
	methodName := "DeleteWallet"

	headers := Headers{
		"walletID": id,
	}

	_, err := self.do(
		"DELETE",
		"v1/wallet",
		nil,
		self.GETHeaders(headers),
	)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return nil
}
