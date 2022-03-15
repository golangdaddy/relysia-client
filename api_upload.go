package relysia

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (self *Client) UploadReference(walletID, filename, url, notes string) error {
	methodName := "Upload"

	headers := Headers{
		"walletID": walletID,
	}

	b, _ := json.Marshal(
		map[string]string{
			"type":     "media",
			"fileUrl":  url,
			"filename": filename,
			"notes":    notes,
		},
	)
	println("UPLOAD", string(b))

	_, err := self.do(
		"POST",
		"v1/upload",
		bytes.NewBuffer(b),
		self.POSTHeaders(headers),
	)
	if err != nil {
		return fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return nil
}
