package comment_repo

type Comment struct {
	Id       int64
	NotifyId int64
	AuthorId int64
	Message  string
}
