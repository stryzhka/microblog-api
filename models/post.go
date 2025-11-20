package models

import "github.com/lib/pq"

type Post struct {
	Id          string
	ProfileId   string
	LikesCount  int
	Likes       pq.StringArray
	Content     string
	PicturePath string
	DateCreated string
}

type Like struct {
	ProfileId string
	PostId    string
}
