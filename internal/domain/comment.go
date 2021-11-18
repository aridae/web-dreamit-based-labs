package domain

type Comment struct {
	Id       int64
	NotifyId int64
	AuthorId int64
	Message  string
}

type PostComment struct {
	AuthorId uint64
	NotifyId int64
	Message  string
}
