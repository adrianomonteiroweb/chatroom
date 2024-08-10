package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gorilla/mux"
)

func TestCreateChatRoom(t *testing.T) {
    req, err := http.NewRequest("POST", "/chatroom", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/chatroom", createChatRoom).Methods("POST")
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }

    expected := `Chat room created`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}
