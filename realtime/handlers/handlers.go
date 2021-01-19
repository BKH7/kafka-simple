package handlers

import (
	"io"
	"net/http"
	"time"

	"github.com/BKH7/kafka-simple/realtime/jsonify"
	"github.com/BKH7/kafka-simple/realtime/msg"
)

// HealthCheck alive of service
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

// Root Realtime Data from gateway
func Root(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		Greeting(w, r)
	} else if r.Method == http.MethodPost {
		Realtime(w, r)
	}
}

// Greeting service
func Greeting(w http.ResponseWriter, r *http.Request) {
	jsonify.JSON(w)(http.StatusOK, map[string]interface{}{
		"message": "realtime service",
	})
}

// Realtime get data from gateway
func Realtime(w http.ResponseWriter, r *http.Request) {
	var m msg.Msgbody
	jsonify.Bind(r)(&m)

	err := msg.Producer(&m)
	if err != nil {
		jsonify.JSON(w)(http.StatusInternalServerError, err)
		return
	}
	jsonify.JSON(w)(http.StatusCreated, map[string]interface{}{
		"Code":    "OK",
		"Created": time.Now().UTC().Format("2006-01-02T15:04:05Z07:00"),
	})
}
