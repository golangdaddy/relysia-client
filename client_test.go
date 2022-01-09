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

}
