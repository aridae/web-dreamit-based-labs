package apimodels

const (
	FAILURE_MY_ROOM_BOOKING    = "failed to get your rooms: %s"
	FAILURE_ROOM_EVENTS        = "failed to get room events: %s"
	FAILURE_DELETE_EVENT       = "failed to delete room event: %s"
	FAILURE_ADD_EVENT          = "failed to delete room event: %s"
	FAILURE_UPDATE_EVENT       = "failed to delete room event: %s"
	SUCCESS_ROOM_EVENT_DELETED = "successfully deleted room"
)

type Event struct {
	Id     int64  `json:"id"`
	RoomId int64  `json:"roomId"`
	Title  string `json:"title"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Author uint64 `json:"author"`
}

type PostEvent struct {
	RoomId int64  `json:"roomId"`
	Title  string `json:"title"`
	Start  string `json:"start"`
	End    string `json:"end"`
}

type SuccessPostEvent struct {
	Id int64 `json:"id"`
}

type PatchEvent struct {
	RoomId int64  `json:"roomId"`
	Start  string `json:"start"`
	End    string `json:"end"`
}
