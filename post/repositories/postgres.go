package repositories

import (
	"database/sql"
	"fmt"
	"microblog-api/models"
	post2 "microblog-api/post"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(post *models.Post) error {
	_, err := r.db.Exec(`insert into posts (id, profile_id, content, date) values ($1, $2, $3, $4)`, post.Id, post.ProfileId, post.Content, post.DateCreated)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *PostgresRepository) GetById(id string) (*models.Post, error) {
	post := &models.Post{}
	err := r.db.QueryRow(`select id, profile_id, content, date from posts where id = $1`, id).Scan(
		&post.Id,
		&post.ProfileId,
		&post.Content,
		&post.DateCreated,
	)
	if err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return nil, post2.ErrPostNotFound
		}
	}
	return post, err
}

func (r *PostgresRepository) GetByUserId(userId string) []models.Post {
	var posts []models.Post
	rows, err := r.db.Query(`select  id, profile_id, content, date from posts where profile_id = $1`, userId)
	if err != nil {
		fmt.Println(err.Error())
		return posts
	}
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.Id, &post.ProfileId, &post.Content, &post.DateCreated)
		if err != nil {
			fmt.Println(err.Error())
			return posts
		}
		posts = append(posts, *post)
	}
	return posts
}

func (r *PostgresRepository) Delete(userId, id string) error {
	temp, err := r.db.Exec(`delete from posts where profile_id = $1 and id = $2`, userId, id)
	found, err := temp.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
	}
	if found == 0 {
		return post2.ErrPostNotFound
	}
	return nil
}
