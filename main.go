package main

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func main() {
	r := http.NewServeMux()

	r.Handle("/api/v1/conversation", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(V1PostConversationHandler)))

	http.ListenAndServe(":80", handlers.CompressHandler(r))
}
