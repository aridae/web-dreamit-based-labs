package models

type Calendar struct {
	RoomId   int64 `json:"roomId"`
	Len      int8  `json:"len"`
	Calendar []Day `json:"calendar"`
}

type Day struct {
	Number      int8 `json:"number"`
	FullyBooked bool `json:"fullyBooked"`
}

type Event struct {
	Id    int64     `json:"id"`
	Title string    `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
	Author uint64 `json:"author"`
}

type Booking struct {
	RoomId int64 `json:"roomId"`
	Booking Event `json:"event"`
}