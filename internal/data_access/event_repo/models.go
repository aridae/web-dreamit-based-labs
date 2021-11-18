package eventrepo

type Event struct {
	Id       int64
	RoomId   int64
	Title    string
	Start    string
	End      string
	AuthorId uint64
}
