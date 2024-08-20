package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type unauthorized struct {
	Access string  `json:"access"`
	Reason *string `json:"message,omitempty"`
}

func Unauthorized(rw http.ResponseWriter, message *string) {
	content := unauthorized{
		Access: "unauthorized",
	}
	if message != nil {
		content.Reason = message
	}
	writeError := makeJsonResponse(rw, http.StatusUnauthorized, content)
	if writeError != nil {
		log.Println("Error writing error message: ", writeError.Error())
	}
	log.Println("Unauthorized: ", message)
}

func makeJsonResponse(rw http.ResponseWriter, status int, data any) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	_, err = rw.Write(js)
	if err != nil {
		return err
	}
	return nil
}
