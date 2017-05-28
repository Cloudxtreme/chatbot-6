package main

import (
	"encoding/json"
	"net/http"
)

func V1PostConversationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var c Conversation
	err := decoder.Decode(&c)

	if err != nil {
		panic(err)
	}

	if SaveConversation(&c) == false {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
