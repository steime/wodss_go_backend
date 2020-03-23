package testing

import (
	"bytes"
	handler2 "github.com/steime/wodss_go_backend/handler"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/steime/wodss_go_backend/repositoryImpl"
)

func TestLogin(t *testing.T){
	repo := mySQL.NewMySqlRepository()
	var jsonStr = []byte(`{
		"Email" : "dst1@mail.com",
		"Password" : "SomethingS4fe*"
	}`)
	req, err := http.NewRequest("GET", "/auth/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler2.Login(repo))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
