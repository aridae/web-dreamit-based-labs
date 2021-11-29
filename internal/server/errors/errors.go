package errors

import (
	"github.com/go-park-mail-ru/2021_1_DuckLuck/pkg/errors"
)

func CreateError(err error) error {
	return errors.CreateError(err)
}

var (
	ErrRoomNotFound error = errors.Error{
		Message: "room not found",
	}
	ErrBadTime error = errors.Error{
		Message: "bad time",
	}
	ErrEventNotFound error = errors.Error{
		Message: "event not found",
	}
)
