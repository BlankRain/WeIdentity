package entity

type PublicKeyProperty struct {
	Id        string
	PkType    string `json:"type"`
	Owner     string
	PublicKey string `json:"publicKey"`
}

type AuthenticationProperty struct {
	PublicKey string `json:"publicKey"`
}
type ServiceProperty struct {
	ServiceEndpoint string `json:"serviceEndpoint"`
}
type long int64 // long in java

type WeIdDocument struct {
	Id             string
	Created        long
	Updated        long
	PublicKey      []PublicKeyProperty
	Authentication []AuthenticationProperty
	Service        []ServiceProperty
}
