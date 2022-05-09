package relysia

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/kr/pretty"
)

type Client struct {
	httpClient *http.Client
	host       string
	authToken  string
	serviceID  string
}

func NewClient() *Client {

	client := &Client{
		host: "https://api.relysia.com",
		httpClient: &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 5 * time.Second,
			},
		},
	}
	return client
}

func (self *Client) WithToken(token string) *Client {
	newClient := *self
	newClient.authToken = token
	return &newClient
}

func (self *Client) WithService(id string) *Client {
	newClient := *self
	newClient.serviceID = id
	return &newClient
}

func (self *Client) SetAccessToken(token string) {
	self.authToken = token
}

func (self *Client) do(method, path string, x io.Reader, headers Headers) ([]byte, error) {

	url := fmt.Sprintf("%s/%s", self.host, path)
	println(url)

	req, err := http.NewRequest(
		method,
		url,
		x,
	)
	if err != nil {
		return nil, fmt.Errorf("DoRequest: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
		//println("HEADER", k, v)
	}

	resp, err := self.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("DoRequest: %w", err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("DoRequest: %w", err)
	}
	resp.Body.Close()

	response := &Response{}
	if len(b) > 0 {
		if err := json.Unmarshal(b, response); err != nil {
			return nil, fmt.Errorf("DoRequest: Invalid response from server: %w", err)
		}
	}
	if resp.StatusCode != 200 || response.StatusCode != 200 {
		var unk interface{} = new(interface{})
		log.Println(string(b))
		json.Unmarshal(b, unk)
		pretty.Println(unk)
		return nil, fmt.Errorf("Invalid response status code from %s: %d: %s", self.host, resp.StatusCode, response.Message)
	}

	return []byte(response.Data), nil
}
