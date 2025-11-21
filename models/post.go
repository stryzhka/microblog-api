package models

type Post struct {
	Id          string
	ProfileId   string
	LikesCount  int
	Likes       []string
	Content     string
	PicturePath string
	DateCreated string
}

type Like struct {
	ProfileId string
	PostId    string
}
