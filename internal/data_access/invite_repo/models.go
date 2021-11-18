package inviterepo

// модели уровня бд

// на стороне БД это справочник
type InviteStatus struct {
	Id     uint64
	Status string
}

type Invite struct {
	Id         int64
	EventId    int64
	ReceiverId uint64
	StatusId   uint64
}
