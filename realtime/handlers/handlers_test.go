package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BKH7/kafka-simple/realtime/handlers"
	"github.com/BKH7/kafka-simple/realtime/msg"
	"gopkg.in/go-playground/assert.v1"
)

func TestGETRoot(t *testing.T) {
	t.Run("it should return httpCode 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.Root)

		handler.ServeHTTP(resp, req)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("it should service response alive", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/health-check", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.HealthCheck)
		handler.ServeHTTP(resp, req)

		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `{"alive": true}`
		if resp.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				resp.Body.String(), expected)
		}
	})
}

func TestPOSTRoot(t *testing.T) {
	body := msg.Msgbody{
		ID:     1,
		Sender: "Tom",
		Msg:    "Hello",
	}
	byte, _ := json.Marshal(body)
	reqReader := bytes.NewReader(byte)

	r, _ := http.NewRequest(http.MethodPost, "/", reqReader)
	w := httptest.NewRecorder()
	handlers.Root(w, r)

	t.Run("should be return 201", func(t *testing.T) {
		got := w.Code
		expected := 201
		assert.Equal(t, expected, got)
	})
}
