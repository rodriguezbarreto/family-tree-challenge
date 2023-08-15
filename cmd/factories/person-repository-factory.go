package factories

import (
	"family-tree-challenge/internal/infra/database"
	"family-tree-challenge/internal/infra/repositories"
)

func PersonRespositoryFactory() *repositories.PersonRespository {
	connection := database.PostgresConnection()
	repository := repositories.NewPersonRepository(connection)

	return repository
}
