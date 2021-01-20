package jsonify

import (
	"encoding/json"
	"net/http"
)

//Bind func has get data from (r *http.Request)
func Bind(r *http.Request) func(interface{}) error {
	return func(v interface{}) error {
		return json.NewDecoder(r.Body).Decode(v)
	}
}

// JSON has func to return data to (w http.ResponseWriter)
func JSON(w http.ResponseWriter) func(int, interface{}) error {
	return func(code int, v interface{}) error {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(code)
		return json.NewEncoder(w).Encode(v)
	}
}
