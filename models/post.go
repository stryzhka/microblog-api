package models

type Post struct {
	Id          string
	ProfileId   string
	IsComment   bool
	LikesCount  int
	Likes       []string
	Content     string
	PicturePath string
	DateCreated string
	Comments    []string
}

type Like struct {
	ProfileId string
	PostId    string
}

type CommentData struct {
	PostId    string
	CommentId string
}
