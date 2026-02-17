package request_handler

import (
	"encoding/json"
	"net/http"
)

var contentType = "Content-Type"
var applicationJson = "application/json"

func ResponseByJson(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set(contentType, applicationJson)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
