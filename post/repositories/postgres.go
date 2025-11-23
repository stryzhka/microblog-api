package repositories

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
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
	_, err := r.db.Exec(`insert into posts (id, profile_id, content, date, picture_path, likes_count, likes, comments, is_comment) values ($1, $2, $3, $4, $5, 0, $6, $7, false)`, post.Id, post.ProfileId, post.Content, post.DateCreated, post.PicturePath, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *PostgresRepository) GetById(id string) (*models.Post, error) {
	post := &models.Post{}
	err := r.db.QueryRow(`select id, profile_id, content, date, likes_count, picture_path , likes, comments, is_comment from posts where id = $1`, id).Scan(
		&post.Id,
		&post.ProfileId,
		&post.Content,
		&post.DateCreated,
		&post.LikesCount,
		&post.PicturePath,
		pq.Array(&post.Likes),
		pq.Array(&post.Comments),
		&post.IsComment,
	)
	if err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return nil, post2.ErrPostNotFound
		}
	}
	return post, err
}

func (r *PostgresRepository) GetAll() []models.Post {
	var posts []models.Post
	rows, err := r.db.Query(`select  id, profile_id, content, date, likes_count, picture_path, likes, comments, is_comment from posts  where is_comment <> true`)
	if err != nil {
		fmt.Println(err.Error())
		return posts
	}
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.Id, &post.ProfileId, &post.Content, &post.DateCreated, &post.LikesCount, &post.PicturePath, pq.Array(&post.Likes), pq.Array(&post.Comments), &post.IsComment)
		if err != nil {
			fmt.Println(err.Error())
			return posts
		}
		posts = append(posts, *post)
	}
	return posts
}

func (r *PostgresRepository) GetByUserId(userId string) []models.Post {
	var posts []models.Post
	rows, err := r.db.Query(`select  id, profile_id, content, date, likes_count, picture_path, likes, comments, is_comment from posts where profile_id = $1`, userId)
	if err != nil {
		fmt.Println(err.Error())
		return posts
	}
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.Id, &post.ProfileId, &post.Content, &post.DateCreated, &post.LikesCount, &post.PicturePath, pq.Array(&post.Likes), pq.Array(&post.Comments), &post.IsComment)
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

func (r *PostgresRepository) LikePost(like *models.Like) error {
	var postId int
	err := r.db.QueryRow(`select id from posts where id = $1`, like.PostId).Scan(&postId)
	if err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return post2.ErrPostNotFound
		}
	}
	_, err = r.db.Exec(`insert into likes (profile_id, post_id) values ($1, $2)`, like.ProfileId, like.PostId)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			if e.Code == "23505" {
				return post2.ErrAlreadyLiked
			}
		}
	}
	_, err = r.db.Exec(`
		UPDATE posts
		SET likes_count = (
			SELECT COUNT(*)
			FROM likes
			WHERE likes.post_id = posts.id
		),
		likes = array_append(likes, $2)
		WHERE id = $1;

	`, like.PostId, like.ProfileId)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) DislikePost(like *models.Like) error {
	var postId int
	err := r.db.QueryRow(`select id from posts where id = $1`, like.PostId).Scan(&postId)
	if err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return post2.ErrPostNotFound
		}
	}
	fmt.Println(like.PostId, like.ProfileId)
	temp, err := r.db.Exec(`delete from likes where profile_id = $1 and post_id = $2`, like.ProfileId, like.PostId)
	found, err := temp.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if found == 0 {
		return post2.ErrNotLiked
	}
	_, err = r.db.Exec(`
		UPDATE posts
		SET likes_count = (
			SELECT COUNT(*)
			FROM likes
			WHERE likes.post_id = posts.id
		),
		likes = array_remove(likes, $2)
		WHERE id = $1;

	`, like.PostId, like.ProfileId)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) AddComment(comment *models.Post, commentData *models.CommentData) error {

	_, err := r.db.Exec(`insert into posts (id, profile_id, content, date, picture_path, likes_count, likes, is_comment) values ($1, $2, $3, $4, $5, 0, $6, $7)`, comment.Id, comment.ProfileId, comment.Content, comment.DateCreated, comment.PicturePath, nil, true)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = r.db.Exec(`insert into comments (post_id, comment_id) values ($1, $2)`, commentData.PostId, commentData.CommentId)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = r.db.Exec(`
		UPDATE posts 
		SET comments = array_append(comments, $1)
		WHERE id = $2;

	`, commentData.CommentId, commentData.PostId)
	if err != nil {
		return err
	}
	return nil
}
