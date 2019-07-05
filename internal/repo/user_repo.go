package repo

import "github.com/stephram/nsm-userdata-service/restapi/models"

type UserRepo interface {
	FindUser(tokenID string) (*models.User, error)
	StoreUser(data *models.User) (*models.User, error)

	FindGameOnResults(tokenID string) (*models.GameOnResults, error)
	StoreGameOnResults(tokenID string, results *models.GameOnResults) (*models.GameOnResults, error)
}
