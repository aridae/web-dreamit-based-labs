package apimodels

const (
	PENDING_STATUS  = iota
	ACCEPTED_STATUS = iota
	DECLINED_STATUS = iota
	EXPIRED_STATUS  = iota
)

type Invite struct {
	Id         int64  `json:"id"`
	EventId    int64  `json:"eventId"`
	ReceiverId uint64 `json:"receiverId"`
	StatusId   uint64 `json:"status"`
}

type PostInvite struct {
	EventId    int64  `json:"eventId"`
	ReceiverId uint64 `json:"receiverId"`
}

type SuccessPostInvite struct {
	Id int64 `json:"id"`
}

type PatchInvite struct {
	StatusId uint64 `json:"status"`
}
