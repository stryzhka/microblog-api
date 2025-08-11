package models

type Post struct {
	Id          string
	UserId      string
	Likes       int
	Content     string
	PicturePath string
	DateCreated string
}
