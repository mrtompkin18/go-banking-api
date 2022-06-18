package errs

import (
	"encoding/json"
	"net/http"
)

type ServiceErrors struct {
	HttpStatus int    `json:"http_status"`
	Message    string `json:"messages"`
}

func (se ServiceErrors) Error() string {
	return se.Message
}

func NewNotFoundError(message string) ServiceErrors {
	return ServiceErrors{
		HttpStatus: http.StatusNotFound,
		Message:    message,
	}
}

func NewInternalError(message string) ServiceErrors {
	return ServiceErrors{
		HttpStatus: http.StatusInternalServerError,
		Message:    message,
	}
}

func NewValidateError(message string) ServiceErrors {
	return ServiceErrors{
		HttpStatus: http.StatusUnprocessableEntity,
		Message:    message,
	}
}

func HandleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case ServiceErrors:
		e.WriteResponse(w)
		break
	default:
		NewInternalError(e.Error()).WriteResponse(w)
		break
	}
}

func (se ServiceErrors) WriteResponse(w http.ResponseWriter) {
	w.WriteHeader(se.HttpStatus)
	err := json.NewEncoder(w).Encode(se)
	if err != nil {
		panic(err)
	}
}
