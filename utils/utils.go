package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

// WriteError sends http error with json body
func WriteError(w http.ResponseWriter, code int, msg string) {
	log.Println(msg)
	WriteJSON(w, code, map[string]string{"error": msg})
}

// WriteJSON sends http response with a json content
func WriteJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// WriteBlob sends http response with a media content
func WriteBlob(w http.ResponseWriter, code int, reader io.Reader, length int64, mimeType string) {
	w.Header().Set("Content-Type", mimeType)
	w.Header().Set("Content-Length", strconv.FormatInt(length, 10))

	if _, err := io.Copy(w, reader); err != nil {
		log.Println("unable to write blob.")
	}
}
