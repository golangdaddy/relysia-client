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
	TokenSupply  int    `json:"tokenSupply,omitempty"`
	TotalSupply  int    `json:"totalSupply,omitempty"`
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
			Media []*MetaMedia `json:"media"`
		} `json:"meta"`
	} `json:"properties"`
	Splitable bool `json:"splitable,omitempty"`
	//
	ContractTxid    string `json:"contractTxid,omitempty"`
	IssueTxid       string `json:"issueTxid,omitempty"`
	IntialSupply    int    `json:"intialSupply,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
	CreationDate    string `json:"creationDate,omitempty"`
	UserID          string `json:"userId,omitempty"`
}

func (self *IssueRequest) ToJSON() []byte {
	b, _ := json.Marshal(self)
	return b
}

type MetaMedia struct {
	URI    string `json:"URI"`
	Type   string `json:"type"`
	AltURI string `json:"altURI"`
}

type IssueResponse struct {
	Status   string       `json:"status"`
	Msg      string       `json:"msg"`
	TokenID  string       `json:"tokenId"`
	TokenObj IssueRequest `json:"tokenObj"`
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

	fmt.Printf("TOKEN RESPONSE %x \n", b)

	response := &IssueResponse{}
	if err := json.Unmarshal(b, response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response, nil
}

func (self *Client) GetToken(id string) (map[string]interface{}, error) {
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

	response := map[string]interface{}{}
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, fmt.Errorf("%s: %s", methodName, err.Error())
	}

	return response, nil
}

func DemoTokenRequest() *IssueRequest {
	issueRequest := &IssueRequest{}
	issueRequest.ProtocolID = "STAS"
	issueRequest.Name = "GO CLIENT TEST TOKEN"
	issueRequest.Image = "https://picsum.photos/200"
	issueRequest.Description = "a test token for github.com/golangdaddy/relysia-client"
	issueRequest.Symbol = "TST"
	issueRequest.TokenSupply = 1
	issueRequest.Decimals = 0
	issueRequest.Splitable = false
	issueRequest.SatsPerToken = 1500
	issueRequest.Properties.Legal.LicenceID = "?"
	issueRequest.Properties.Legal.Terms = "?"
	issueRequest.Properties.Issuer.Email = "?"
	issueRequest.Properties.Issuer.LegalForm = "?"
	issueRequest.Properties.Issuer.Organisation = "?"
	issueRequest.Properties.Issuer.IssuerCountry = "?"
	issueRequest.Properties.Issuer.GoverningLaw = "?"
	issueRequest.Properties.Issuer.Jurisdiction = "?"
	issueRequest.Properties.Meta.SchemaID = "NFT1.0/MA"
	issueRequest.Properties.Meta.Website = "vaionex.com"
	issueRequest.Properties.Meta.Legal.Terms = "© 2020 TAAL TECHNOLOGIES SEZC\nALL RIGHTS RESERVED. ANY USE OF THIS SOFTWARE IS SUBJECT TO TERMS AND CONDITIONS OF LICENSE. USE OF THIS SOFTWARE WITHOUT LICENSE CONSTITUTES INFRINGEMENT OF INTELLECTUAL PROPERTY. FOR LICENSE DETAILS OF THE SOFTWARE, PLEASE REFER TO: www.taal.com/stas-token-license-agreement"
	issueRequest.Properties.Meta.Media = append(
		issueRequest.Properties.Meta.Media,
		&MetaMedia{
			URI:    "https://picsum.photos/200",
			Type:   "image",
			AltURI: "https://picsum.photos/200",
		},
	)
	return issueRequest
}
