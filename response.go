package relysia

import "encoding/json"

type Response struct {
	StatusCode int             `json:"statusCode"`
	Error      string          `json:"error"`
	Message    string          `json:"message,omitempty"`
	Data       json.RawMessage `json:"data"`
}
