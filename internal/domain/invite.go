package domain

const (
	PENDING_STATUS  = iota
	ACCEPTED_STATUS = iota
	DECLINED_STATUS = iota
	EXPIRED_STATUS  = iota
)

type Invite struct {
	Id         int64
	EventId    int64
	ReceiverId uint64
	StatusId   uint64
}

type PostInvite struct {
	EventId    int64
	ReceiverId uint64
}
