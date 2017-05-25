package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
    host = "",
    port = "8080"
)

func PostConversationEndPoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello!"))
}

func main() {
	router := httprouter.New()
	router.POST("/", PostConversationEndPoint)

    log.Fatal(http.ListenAndServe(host + ":" port, router))
}
