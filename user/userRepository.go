package user

import "microblog-api/models"

type Repository interface {
	//List() (models.User, error)
	Create(user *models.User) error
	//GetById(id string) (models.User, error)
	//GetByUsername(username string) (models.User, error)
	Get(username, password string) (*models.User, error)
	//Delete(id string) error
	//Update(id string) (models.User, error)
}
