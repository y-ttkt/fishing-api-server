package errors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Code    ErrCode  `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func Handler(w http.ResponseWriter, err error) {
	var httpErr HTTPError
	if !errors.As(err, &httpErr) {
		log.Printf("Unhandled error: %v", err)
		httpErr = NewAppError(Unknown, err.Error(), err)
	}

	log.Printf("AppError occurred: Code=%s, Message=%s, Error=%v",
		httpErr.Code(),
		httpErr.Message(),
		err,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErr.Status())
	json.NewEncoder(w).Encode(ErrorResponse{
		Code:    httpErr.Code(),
		Message: httpErr.Message(),
		Errors:  []string{httpErr.Error()},
	})
}
