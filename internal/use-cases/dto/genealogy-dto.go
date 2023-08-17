package dto

type Genealogy struct {
	Members []Member `json:"members"`
}

type Member struct {
	Name          string         `json:"name"`
	Relationships []Relationship `json:"relationships"`
}

type Relationship struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
}
