package errors

import (
	"net/http"
)

type ErrCode string

const (
	Unknown ErrCode = "U000"

	NotFound     ErrCode = "S001"
	Invalid      ErrCode = "S002"
	Conflict     ErrCode = "S003"
	BadRequest   ErrCode = "S004"
	Unauthorized ErrCode = "S005"
	InternalErr  ErrCode = "E100"
)

var errStatusMap = map[ErrCode]int{
	NotFound:     http.StatusNotFound,
	Invalid:      http.StatusUnprocessableEntity,
	Conflict:     http.StatusConflict,
	BadRequest:   http.StatusBadRequest,
	Unauthorized: http.StatusUnauthorized,
	InternalErr:  http.StatusInternalServerError,

	// 新しいエラーコードを追加する場合はここにエントリを追加
}

func (code ErrCode) DefaultHTTPStatus() int {
	if status, ok := errStatusMap[code]; ok {
		return status
	}

	return http.StatusInternalServerError
}

func (code ErrCode) Wrap(message string, err error) *AppError {
	return NewAppError(code, message, err)
}
