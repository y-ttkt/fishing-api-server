package errors

import (
	"net/http"
)

type ErrCode string

const (
	Unknown ErrCode = "U000"

	NotFound     ErrCode = "S001"
	InsertFailed ErrCode = "S002"
	Invalid      ErrCode = "S003"
	SelectFailed ErrCode = "S004"
	Conflict     ErrCode = "S005"
)

var errStatusMap = map[ErrCode]int{
	NotFound:     http.StatusNotFound,
	Invalid:      http.StatusUnprocessableEntity,
	Conflict:     http.StatusConflict,
	InsertFailed: http.StatusInternalServerError,
	SelectFailed: http.StatusInternalServerError,

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
