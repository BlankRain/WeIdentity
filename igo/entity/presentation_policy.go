package entity

type ClaimPolicy struct {
	FieldsToBeDisclosed string
}
type PresentationPolicy struct {
	Id                  int64
	OrgId               string
	Version             int32
	PolicyPublisherWeId string
	Policy              map[int64]ClaimPolicy
	Extra               map[string]string
}
