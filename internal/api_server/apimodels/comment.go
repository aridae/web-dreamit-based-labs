package apimodels

type Comment struct {
	Id       int64  `json:"id"`
	NotifyId int64  `json:"notifyId"`
	AuthorId int64  `json:"authorId"`
	Message  string `json:"message"`
}

type PostComment struct {
	NotifyId int64  `json:"notifyId"`
	Message  string `json:"message"`
}

type SuccessPostComment struct {
	Id int64 `json:"id"`
}
