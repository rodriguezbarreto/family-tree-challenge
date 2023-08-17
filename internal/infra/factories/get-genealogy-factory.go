package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func GetGenealogyFactory() *controllers.GetGenealogyController {
	repoRel := RelationshipRespositoryFactory()
	repoPerson := PersonRespositoryFactory()
	GetGenealogyUseCase := usecases.NewGetGenealogy(repoRel, repoPerson)
	GetGenealogyController := controllers.NewGetGenealogyController(GetGenealogyUseCase)

	return GetGenealogyController
}
