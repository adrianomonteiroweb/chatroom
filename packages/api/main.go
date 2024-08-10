package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func createChatRoom(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Chat room created"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/chatroom", createChatRoom).Methods("POST")

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
