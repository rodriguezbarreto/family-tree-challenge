package usecases

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/dto"
	"family-tree-challenge/internal/use-cases/repositories"
)

type GetBaconNumber struct {
	repoRel    repositories.RelationshipRepository
	repoPerson repositories.PersonRepository
	allRels    []*domain.Relationship
	allPersons []*domain.Person
}

func NewGetBaconNumber(repoRel repositories.RelationshipRepository, repoPerson repositories.PersonRepository) *GetBaconNumber {
	return &GetBaconNumber{
		repoRel:    repoRel,
		repoPerson: repoPerson,
	}
}

func (u *GetBaconNumber) Execute(sourceID, targetID string) (int, error) {
	err := u.initialActions()
	if err != nil {
		return -1, err
	}

	graph := u.buildGraph()
	distance := u.bfs(graph, sourceID, targetID)
	if distance == -1 {
		return -1, err
	}

	return distance, nil
}

func (u *GetBaconNumber) buildGraph() map[string][]string {
	graph := make(map[string][]string)
	for _, rel := range u.allRels {
		graph[rel.ParentID] = append(graph[rel.ParentID], rel.ChildID)
		graph[rel.ChildID] = append(graph[rel.ChildID], rel.ParentID)
	}
	return graph
}

func (u *GetBaconNumber) bfs(graph map[string][]string, sourceID, targetID string) int {
	visited := make(map[string]bool)
	queue := []string{sourceID}
	distance := make(map[string]int)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				distance[neighbor] = distance[current] + 1

				if neighbor == targetID {
					return distance[neighbor]
				}
			}
		}
	}

	return -1
}

func (u *GetBaconNumber) initialActions() error {
	allPersons, err := u.repoPerson.List(nil)
	if err != nil {
		return err
	}
	u.allPersons = allPersons

	allRels, err := u.repoRel.List(&dto.RelationshipFilter{})
	if err != nil {
		return err
	}
	u.allRels = allRels

	return nil
}
