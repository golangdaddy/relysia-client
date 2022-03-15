package relysia

import (
	"log"
	"testing"
)

func TestTokenOutput(t *testing.T) {

	log.Println(string(DemoTokenRequest().ToJSON()))

}
