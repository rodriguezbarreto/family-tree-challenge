package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func GetBaconNumberFactory() *controllers.GetBaconNumberController {
	repoRel := RelationshipRespositoryFactory()
	repoPerson := PersonRespositoryFactory()
	getBaconNumberUseCase := usecases.NewGetBaconNumber(repoRel, repoPerson)
	getBaconNumberController := controllers.NewGetBaconNumberController(getBaconNumberUseCase)

	return getBaconNumberController
}
