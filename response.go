package relysia

import "encoding/json"

type Response struct {
	StatusCode int             `json:"statusCode"`
	Error      string          `json:"error"`
	Data       json.RawMessage `json:"data"`
}
