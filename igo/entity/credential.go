package entity

type Credential struct {
	Context        string
	CredentialType string `json:"type"`
	Id             string
	CptId          int64
	Issuer         string
	IssuanceDate   long
	ExpirationDate long
	Claim          map[string]interface{}
	Proof          map[string]interface{}
}
