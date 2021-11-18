package notifyrepo

// повторяет табличку
type Notify struct {
	Id      int64
	EventId int64
	Subject string
	Tags    []string
	Message string
}
