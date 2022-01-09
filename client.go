package relysia

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/kr/pretty"
)

type Client struct {
	httpClient  *http.Client
	host        string
	in          chan *http.Request
	out         chan interface{}
	accessToken string
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
		in:  make(chan *http.Request),
		out: make(chan interface{}),
	}
	go client.rateLimiter()
	return client
}

func (self *Client) SetAccessToken(token string) {
	self.accessToken = token
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
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	self.in <- req

	switch v := (<-self.out).(type) {
	case error:
		return nil, v
	case []byte:
		return v, nil
	case json.RawMessage:
		return []byte(v), nil
	}
	return nil, nil
}

func (self *Client) rateLimiter() {
	for range time.NewTicker(time.Second / 10).C {

		req := <-self.in

		resp, err := self.httpClient.Do(req)
		if err != nil {
			self.out <- err
			continue
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			self.out <- err
			continue
		}
		resp.Body.Close()

		response := &Response{}
		if len(b) > 0 {
			if err := json.Unmarshal(b, response); err != nil {
				self.out <- fmt.Errorf("Invalid response from server: %s", err)
				continue
			}
		}
		if resp.StatusCode != 200 || response.StatusCode != 200 {
			pretty.Println(resp.Status)
			self.out <- fmt.Errorf("Invalid response status code from %s: %d: %s", self.host, resp.StatusCode, response.Error)
			continue
		}
		self.out <- response.Data
	}
}

func (self *Client) InitBeta() error {

	_, err := self.do(
		"GET",
		"v1/initBeta",
		nil,
		map[string]string{
			"serviceID": "",
			"authToken": "",
		},
	)
	if err != nil {
		return fmt.Errorf("DecodeTransaction: %s", err.Error())
	}

	return nil
}
