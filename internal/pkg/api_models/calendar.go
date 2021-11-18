package api_models

type Event struct {
	Id     int64  `json:"id"`
	RoomId int64  `json:"roomId"`
	Title  string `json:"title"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Author uint64 `json:"author"`
}
