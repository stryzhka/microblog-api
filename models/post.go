package models

type Post struct {
	Id          string
	ProfileId   string
	Likes       int
	Content     string
	PicturePath string
	DateCreated string
}

type Like struct {
	ProfileId string
	PostId    string
}
