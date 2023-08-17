package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func GetGenealogyFactory() *controllers.GetGenealogyController {
	repoRel := RelationshipRespositoryFactory()
	repoPerson := PersonRespositoryFactory()
	getGenealogyUseCase := usecases.NewGetGenealogy(repoRel, repoPerson)
	getGenealogyController := controllers.NewGetGenealogyController(getGenealogyUseCase)

	return getGenealogyController
}
