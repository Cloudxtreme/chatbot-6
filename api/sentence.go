package main

import (
	"log"
)

type Sentence struct {
	Sentence string `json:sentence`
}

type Sentences []Sentence

func SaveSentence(sentence *Sentence) bool {
	session := DBSession()
	defer session.Close()

	if err := session.Query(`INSERT INTO sentences(sentence) VALUES (?)`,
		sentence.Sentence).Exec(); err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
