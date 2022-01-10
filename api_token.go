package relysia

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type IssueRequest struct {
	Name         string `json:"name"`
	ProtocolID   string `json:"protocolId"`
	Symbol       string `json:"symbol"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	TokenSupply  int    `json:"tokenSupply"`
	Decimals     int    `json:"decimals"`
	SatsPerToken int    `json:"satsPerToken"`
	Properties   struct {
		Legal struct {
			Terms     string `json:"terms"`
			LicenceID string `json:"licenceId"`
		} `json:"legal"`
		Issuer struct {
			Organisation  string `json:"organisation"`
			LegalForm     string `json:"legalForm"`
			GoverningLaw  string `json:"governingLaw"`
			IssuerCountry string `json:"issuerCountry"`
			Jurisdiction  string `json:"jurisdiction"`
			Email         string `json:"email"`
		} `json:"issuer"`
		Meta struct {
			SchemaID string `json:"schemaId"`
			Website  string `json:"website"`
			Legal    struct {
				Terms string `json:"terms"`
			} `json:"legal"`
			Media []struct {
				URI    string `json:"URI"`
				Type   string `json:"type"`
				AltURI string `json:"altURI"`
			} `json:"media"`
		} `json:"meta"`
	} `json:"properties"`
	Splitable bool `json:"splitable"`
}

func (self *Client) Issue(headers Headers, issueRequest *IssueRequest) (float64, error) {
	methodName := "Issue"

	b, err := json.Marshal(issueRequest)
	if err != nil {
		return 0, fmt.Errorf("%s: %s", methodName, err.Error())
	}
	println(string(b))

	b, err = self.do(
		"POST",
		"v1/issue",
		bytes.NewBuffer(b),
		self.POSTHeaders(headers),
	)
	if err != nil {
		return 0, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &CurrencyResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return 0, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response.Balance, nil
}
