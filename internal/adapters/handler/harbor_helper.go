package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// writeSuccessResponse will avoid repetition of the code used for create a
// default
func writeSuccessResponse(w http.ResponseWriter, code int, id string) http.ResponseWriter {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write(newCreatedResponse(code, id))
	if err != nil {
		log.Printf("failed to write response body: %v", err)
	}
	return w
}

func writeErrorResponse(w http.ResponseWriter, code int, reason string) http.ResponseWriter {
	log.Print("failed to unmarshal request body")
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write(newJSONErrorBody(code, reason))
	if err != nil {
		log.Printf("failed to write error response body: %v", err)
	}
	return w
}

func newCreatedResponse(code int, id string) []byte {
	log.Print("successfully stored new harbor in database")
	r, err := json.Marshal(map[string]string{"ID": id, "result": http.StatusText(code)})
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	return r
}

func newJSONErrorBody(code int, reason string) []byte {
	r, err := json.Marshal(
		struct {
			Code   int    `json:"http_code"`
			Err    string `json:"error"`
			Reason string `json:"reason"`
		}{
			Code:   code,
			Err:    http.StatusText(code),
			Reason: reason,
		},
	)
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	return r
}
