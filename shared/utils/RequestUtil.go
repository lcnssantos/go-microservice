package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ThrowHttpError(w http.ResponseWriter, status int, message string) {
	answer, _ := json.Marshal(HttpError{Message: message, Status: status})
	w.WriteHeader(status)
	w.Write(answer)
	return
}

func HandleValidateRequest(w http.ResponseWriter, r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		ThrowHttpError(w, http.StatusBadRequest, "Invalid body request")
		return err
	}

	if err := validator.New().Struct(data); err != nil {
		ThrowHttpError(w, http.StatusBadRequest, err.Error())
		return err
	}

	return nil
}
