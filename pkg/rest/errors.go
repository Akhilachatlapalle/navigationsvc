package rest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"runtime"
)

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err := w.Write(buf.Bytes())
	HandleError(err)
}

// HandleError logs errors, used to handleError on defer
func HandleError(err error) {
	if err != nil {
		stack := make([]byte, 4<<10)          // 4 KB
		length := runtime.Stack(stack, false) // only current goroutine
		log.Printf("HandleError: %v: %s", err, string(stack[:length]))
	}
}

type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewError(statusCode int, msg string) Error {
	return Error{
		StatusCode: statusCode,
		Message:    msg,
	}
}
