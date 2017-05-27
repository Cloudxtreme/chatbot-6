package main

import (
	"log"
)

type Conversation struct {
	Question string `json:question`
	Response string `json:response`
}

type Conversations []Conversation

func SaveConversation(conversation *Conversation) bool {
	session := DBSession()
	defer session.Close()

    // check if question/response combination already exists
    var hits int
    if err := session.Query(`SELECT hits FROM conversations WHERE question = ? AND response = ? LIMIT 1`, conversation.Question, conversation.Response).Scan(&hits); err != nil {
        hits = 0
    }

    // save conversation
	if err := session.Query(`INSERT INTO conversations(question, response, hits) VALUES (?, ?, ?)`,
		conversation.Question, conversation.Response, hits + 1).Exec(); err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
