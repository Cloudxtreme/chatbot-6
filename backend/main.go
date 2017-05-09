package main

import(
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type Message struct {
    Message string `json:"message,omitempty"`
}

func GetMessageEnpoint(w http.ResponseWritter, req *http.Request) {
    params := mux.Vars(req)
    json.NewEncoder(w).Encode("Lorem Ipsum")
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", GetMessageEnpoint).Methods("GET")
    log.Fatal(http.listenAndServer(":8080", router))
}
