package controllers

import (
	"encoding/json"
	"net/http"
)

// PingControllerHandler will hold everything that controller needs
type PingControllerHandler struct {
}

// NewPingController returns an instance of the controller handler
func NewPingController() *PingControllerHandler {
	return &PingControllerHandler{}
}

// Ping returns http 200 if build is okay and able to serve
func (handler *PingControllerHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(`{"status":"ok"}`), &response); err != nil {
		panic(error(err))
	}
	json.NewEncoder(w).Encode(response)
}
