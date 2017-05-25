package main

import (
	"log"
)

type Conversation struct {
	Question string `json:question`
	Response string `json:response`
	Hits     int    `json:hits`
}

type Conversations []Conversation

func SaveConversation(conversation *Conversation) bool {
	session := DBSession()
	defer session.Close()

	if err := session.Query(`INSERT INTO conversations(question, response, hits) VALUES (?, ?, ?)`,
		conversation.Question, conversation.Response, 1).Exec(); err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
