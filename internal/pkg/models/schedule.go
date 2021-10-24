package models

import "time"

type Schedule struct {
	RoomId   int64       `json:"roomId"`
	Len      int8        `json:"len"`
	Schedule []*Interval `json:"schedule"`
}

type Interval struct {
	Id          int8         `json:"number"`
	IsBooked    bool         `json:"isBooked"`
	BookingTime *BookingTime `json:"bookingTime"`
}

type BookingTime struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type UpdateBooking struct {
	IsBooked  int64  `json:"isBooked"`
	Intervals []int8 `json:"intervals"`
}
