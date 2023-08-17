package dto

type RelationshipInputDTO struct {
	Child  string `json:"child"`
	Parent string `json:"parent"`
}

type RelationshipFilter struct {
	RelID    *string `json:"relId"`
	ChildID  *string `json:"childId"`
	ParentID *string `json:"parentId"`
}

type BaconNumberInputDTO struct {
	SourceID string `json:"sourceId"`
	TargetID string `json:"targetId"`
}
