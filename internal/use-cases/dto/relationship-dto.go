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
