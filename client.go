package relysia

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	host       string
	authToken  string
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
		return nil, fmt.Errorf("DoRequest: %s", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
		//println("HEADER", k, v)
	}

	resp, err := self.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("DoRequest: %s", err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("DoRequest: %s", err)
	}
	resp.Body.Close()

	response := &Response{}
	if len(b) > 0 {
		if err := json.Unmarshal(b, response); err != nil {
			return nil, fmt.Errorf("DoRequest: Invalid response from server: %s", err)
		}
	}
	if resp.StatusCode != 200 || response.StatusCode != 200 {
		println(string(b))
		return nil, fmt.Errorf("Invalid response status code from %s: %d: %s", self.host, resp.StatusCode, response.Message)
	}

	return []byte(response.Data), nil
}

func (self *Client) InitBeta() error {
	methodName := "InitBeta"

	_, err := self.do(
		"GET",
		"v1/initBeta",
		nil,
		self.GETHeaders(),
	)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return nil
}
