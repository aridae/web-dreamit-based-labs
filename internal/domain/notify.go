package domain

// ВОПРОС:
// каскад автоматизированных действий на уровне бд реализуется триггерами
// например: удалили ивент, все приглосы становятся expired
// или, создали нотифай для ивента, отправили пгдемону сигнал дернуть скрипт
// и разослали всем приглашенным уведомления (нотифай) на почту

// а если не на уровне бд?
// типа создали нотифай - надо разослать уведомления на почту
// для этого на сервере есть почтовый виртуальный сервер - демон-почтальон
// это считается внешней сущностью?

type Notify struct {
	Id      int64
	Tags    []string
	Subject string
	EventId int64
	Message string
}

type PostNotify struct {
	Tags    []string
	Subject string
	EventId int64
	Message string
}

// нужно отфильтровать коллекцию нотифаев по параметрам
// прописывать все возможные комбинации параметров для
// разных юскейсов больно (их много),
// поэтому заводим фильтр для коллекции
// этот фильтр передается в репозиторий для селекта
type RequiredNotifyFilter struct {
	EventId int64
}

type OptionalNotifyFilter struct {
	RequiredNotifyFilter // embed struct
	Tags                 []string
	Subject              string
}
