package repositories

import (
	"database/sql"
	"github.com/lib/pq"
	"microblog-api/auth"
	"microblog-api/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(user *models.User) error {
	_, err := r.db.Exec(`INSERT INTO users (id, role, username, password) VALUES ($1, $2, $3, $4)`,
		user.Id, user.Role, user.Username, user.Password)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			if e.Code == "23505" {
				return auth.ErrUserAlreadyExists
			}
		}
	}
	return err
}

func (r *PostgresRepository) Get(username, password string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(`
		SELECT id, username, password, role 
		FROM users 
		WHERE username = $1 AND password = $2`, username, password).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			if e.Code == "23505" {
				return nil, auth.ErrUserAlreadyExists
			}
		}
	}
	return user, err
}

func (r *PostgresRepository) GetUserId(username string) (string, error) {
	user := &models.User{}
	err := r.db.QueryRow(`
		SELECT id 
		FROM users 
		WHERE username = $1 AND`, username).Scan(
		&user.Id,
	)

	return user.Id, err
}
