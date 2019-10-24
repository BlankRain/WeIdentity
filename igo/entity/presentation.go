package entity

type Presentation struct {
	Context          []string
	PresentationType string `json:"type"`
	CredentialList   []Credential
	Proof            map[string]interface{}
}
