package relysia

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	EXAMPLE_IMAGE_URL = "https://www.relysia.com/_next/image?url=%2F_next%2Fstatic%2Fmedia%2FRelysiaLogo_1.5da1be3a.svg&w=750&q=75"
)

func TestClient(t *testing.T) {
	assert := assert.New(t)

	client := NewClient()

	email := fmt.Sprintf("%d@cpu.host", time.Now().UTC().Unix())
	password := "this is my password"

	token, err := client.SignUp(email, password)
	assert.Nil(err)
	assert.Greater(len(token), 0)

	token, err = client.Auth(email, password)
	assert.Nil(err)
	assert.Greater(len(token), 0)

	client.SetAccessToken(token)
	println(token)

	assert.Nil(client.FeeMetricsBeta())
	_, err = client.FeeAddressBeta()
	assert.Nil(err)

	userDetails, err := client.User()
	assert.Nil(err)
	assert.NotNil(userDetails)

	defaultWalletID, err := client.CreateWallet("default")
	assert.Nil(err)

	walletList, err := client.Wallets()
	assert.Nil(err)
	assert.Equal(1, len(walletList))

	value, err := client.Balance(defaultWalletID, "STAS", "")
	assert.Nil(err)
	assert.NotNil(value)

	mn, err := client.Mnemonic()
	assert.Nil(err)
	println(mn)
	assert.Greater(len(mn), 0)

	add, pym, err := client.Address(defaultWalletID)
	assert.Nil(err)
	println(add, pym)
	assert.Greater(len(add), 0)
	assert.Greater(len(pym), 0)
	/*
		ads, err := client.AllAddresses("")
		assert.Nil(err)
		assert.Greater(len(ads), 1)
	*/

	hr, err := client.History("")
	assert.Nil(err)
	assert.NotNil(hr)

	oneWalletID, err := client.CreateWallet("one")
	assert.Nil(err)

	walletList, err = client.Wallets()
	assert.Nil(err)
	assert.Equal(2, len(walletList))

	resp, err := client.Issue(
		Headers{
			"walletID":   oneWalletID,
			"protocolID": "stas",
		},
		DemoTokenRequest(),
	)
	assert.Nil(err)
	assert.NotNil(resp)

	value, err = client.Balance(defaultWalletID, "STAS", "")
	assert.Nil(err)
	assert.NotNil(value)

	uploadResponse, err := client.UploadReference(defaultWalletID, "myfile.png", EXAMPLE_IMAGE_URL, "")
	assert.Nil(err)
	assert.NotNil(uploadResponse)
	assert.Len(uploadResponse.UploadObj.URL, 68)

	assert.Nil(client.DeleteWallet(oneWalletID))
	assert.Nil(client.DeleteWallets())
	assert.Nil(client.DeleteUser())

}
