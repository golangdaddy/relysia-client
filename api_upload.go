package relysia

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type UploadResponse struct {
	Status    string `json:"status"`
	Msg       string `json:"msg"`
	UploadObj struct {
		FileName    string    `json:"fileName"`
		FileType    string    `json:"fileType"`
		FileSize    int       `json:"fileSize"`
		TimeStamp   time.Time `json:"timeStamp"`
		Txid        string    `json:"txid"`
		Address     string    `json:"address"`
		AddressPath string    `json:"addressPath"`
		URL         string    `json:"url"`
	} `json:"uploadObj"`
}

func (self *Client) UploadReference(walletID, filename, url, notes string) (*UploadResponse, error) {
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

	b, err := self.do(
		"POST",
		"upload",
		bytes.NewBuffer(b),
		self.POSTHeaders(headers),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	fmt.Println("UPLOADRESP", hex.EncodeToString(b))

	uploadResponse := &UploadResponse{}
	if err := json.Unmarshal(b, uploadResponse); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return uploadResponse, nil
}
