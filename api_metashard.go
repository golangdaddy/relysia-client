package relysia

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (self *Client) Upload() error {
	methodName := "Upload"

	payload := map[string]string{}
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}
	_, err = self.do(
		"POST",
		"v1/upload",
		bytes.NewBuffer(b),
		self.POSTHeaders(),
	)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}
	return nil
}
