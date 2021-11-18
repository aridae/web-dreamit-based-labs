package apimodels

type Notify struct {
	Id      int64    `json:"id"`
	Tags    []string `json:"tags"`
	Subject string   `json:"subject"`
	EventId int64    `json:"eventId"`
	Message string   `json:"message"`
}

type PostNotify struct {
	Tags    []string `json:"tags"`
	Subject string   `json:"subject"`
	EventId int64    `json:"eventId"`
	Message string   `json:"message"`
}

type SuccessPostNotify struct {
	Id int64 `json:"id"`
}
