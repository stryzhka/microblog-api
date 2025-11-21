package auth

import "microblog-api/models"

type Repository interface {
	//List() (models.User, error)
	Create(user *models.User) error
	GetUserId(username string) (string, error)
	//GetByUsername(username string) (models.User, error)
	Get(username, password string) (*models.User, error)
	//Delete(id string) error
	//Update(id string) (models.User, error)
}
