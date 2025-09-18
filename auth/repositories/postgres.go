package repositories

import (
	"database/sql"
	"microblog-api/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) (*PostgresRepository, error) {
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Create(user *models.User) error {
	_, err := r.db.Exec(`INSERT INTO users (id, role, username, password) VALUES ($1, $2, $3)`,
		user.Id, user.Role, user.Password)
	return err
}

func (r *PostgresRepository) Get(username, password string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(`
		SELECT id, username, password, role 
		FROM users 
		WHERE username = $1, password = $2`, username, password).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Role,
	)
	return user, err
}
