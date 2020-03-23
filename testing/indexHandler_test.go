package testing

import (
	handler2 "github.com/steime/wodss_go_backend/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler2.IndexHandler())
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusTeapot {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusTeapot)
	}
}
