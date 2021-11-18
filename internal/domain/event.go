package domain

// модели уровня бизнес логики

type Event struct {
	Id       int64
	RoomId   int64
	Title    string
	Start    string
	End      string
	AuthorId uint64
}

type PostEvent struct {
	RoomId   int64
	Title    string
	Start    string
	End      string
	AuthorId uint64
}

// перенести ивент
type PatchEvent struct {
	Id     int64
	RoomId int64
	Start  string
	End    string
}
