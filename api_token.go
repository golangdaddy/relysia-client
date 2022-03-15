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
			Media []*IR_Media `json:"media"`
		} `json:"meta"`
	} `json:"properties"`
	Splitable bool `json:"splitable"`
}

type IR_Media struct {
	URI    string `json:"URI"`
	Type   string `json:"type"`
	AltURI string `json:"altURI"`
}

type IssueResponse struct {
	StatusCode int `json:"statusCode"`
	Data       struct {
		Status   string `json:"status"`
		Msg      string `json:"msg"`
		TokenID  string `json:"tokenId"`
		TokenObj struct {
			Name         string `json:"name"`
			ProtocolID   string `json:"protocolId"`
			Symbol       string `json:"symbol"`
			Description  string `json:"description"`
			Image        string `json:"image"`
			TotalSupply  int    `json:"totalSupply"`
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
			ContractTxid    string `json:"contractTxid"`
			IssueTxid       string `json:"issueTxid"`
			IntialSupply    int    `json:"intialSupply"`
			ContractAddress string `json:"contractAddress"`
			CreationDate    string `json:"creationDate"`
			UserID          string `json:"userId"`
		} `json:"tokenObj"`
	} `json:"data"`
}

func (self *Client) Issue(headers Headers, issueRequest *IssueRequest) (*IssueResponse, error) {
	methodName := "Issue"

	b, err := json.Marshal(issueRequest)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}
	println(string(b))

	b, err = self.do(
		"POST",
		"v1/issue",
		bytes.NewBuffer(b),
		self.POSTHeaders(headers),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &IssueResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response, nil
}

func (self *Client) GetToken(id string) (*IssueRequest, error) {
	methodName := "GetToken"

	b, err := self.do(
		"GET",
		fmt.Sprintf("v1/token/%s", id),
		nil,
		self.GETHeaders(),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	response := &IssueRequest{}
	if err := json.Unmarshal(b, response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response, nil
}
