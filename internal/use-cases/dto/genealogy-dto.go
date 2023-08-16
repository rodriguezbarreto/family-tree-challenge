package dto

type FamilyMember struct {
	Name          string           `json:"name"`
	Relationships []FamilyRelation `json:"relationships"`
}

type FamilyRelation struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
}