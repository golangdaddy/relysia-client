package relysia

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	assert.Nil(client.CreateWallet("default"))

	walletList, err := client.Wallets()
	assert.Nil(err)
	assert.Equal(1, len(walletList))

	value, err := client.Balance("", "BSV")
	assert.Nil(err)
	assert.NotNil(value)

	mn, err := client.Mnemonic()
	assert.Nil(err)
	println(mn)
	assert.Greater(len(mn), 0)

	add, pym, err := client.Address("")
	assert.Nil(err)
	println(add, pym)
	assert.Greater(len(add), 0)
	assert.Greater(len(pym), 0)
	/*
		ads, err := client.AllAddresses("")
		assert.Nil(err)
		assert.Greater(len(ads), 1)
	*/
	assert.Nil(client.CreateWallet("one"))

	walletList, err = client.Wallets()
	assert.Nil(err)
	assert.Equal(2, len(walletList))

	assert.Nil(client.DeleteWallet("one"))
	assert.Nil(client.DeleteWallets())
	assert.Nil(client.DeleteUser())

}
