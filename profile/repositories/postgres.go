package repositories

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"microblog-api/models"
	profile2 "microblog-api/profile"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(profile *models.Profile) error {
	_, err := r.db.Exec(`INSERT INTO profiles (id, user_id, name, status, photo) VALUES ($1, $2, $3, $4, $5)`,
		profile.Id, profile.UserId, profile.Name, profile.Status, profile.Photo)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			if e.Code == "23505" {
				return profile2.ErrProfileNotFound
			}
		}
	}
	return err
}

func (r *PostgresRepository) GetById(id string) (*models.Profile, error) {
	profile := &models.Profile{}
	err := r.db.QueryRow(`
		SELECT id, user_id, name, status, photo 
		FROM profiles
		WHERE user_id = $1`, id).Scan(
		&profile.Id,
		&profile.UserId,
		&profile.Name,
		&profile.Status,
		&profile.Photo,
	)

	if err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return nil, profile2.ErrProfileNotFound
		}
	}
	return profile, err
}

func (r *PostgresRepository) GetAll() []models.Profile {
	var profiles []models.Profile
	rows, err := r.db.Query(`select user_id, name, status, photo from profiles`)
	if err != nil {
		fmt.Println(err.Error())
		return profiles
	}
	for rows.Next() {
		profile := &models.Profile{}
		err := rows.Scan(&profile.UserId, &profile.Name, &profile.Status, &profile.Photo)
		if err != nil {
			fmt.Println(err.Error())
			return profiles
		}
		profiles = append(profiles, *profile)
	}
	return profiles
}

func (r *PostgresRepository) Update(id, userId string, newProfile *models.Profile) error {
	_, err := r.db.Exec(`
		update profiles set name = $1, status = $2, photo = $3 where user_id = $4
		`, newProfile.Name, newProfile.Status, newProfile.Photo, userId)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			if e.Code == "23505" {
				return profile2.ErrNameAlreadyExists
			}
		}
	}
	return err
}
