package main

import (
    "log"
    "github.com/gocql/gocql"
)

type Question struct {
    Question string `json:question`
    Response string `json:response`
    Hits int `json:hits`
}

type QuestionRepository struct {}

func (q *QuestionRepository) GetAll(question string) ([]map[string]interface{}) {
    session := DBSession()
    defer session.Close()

    var query gocql.session
    query = session.Query(`SELECT question, response, hits FROM questions`)
    if question != "" {
        query = session.Query(`SELECT question, response, hits FROM questions WHERE question = ?`, question)
    }

    iter := query.Iter()
    defer iter.Close()

    questions, err := iter.SliceMap()

    if err != nil {
        log.Fatal(err)
    }

    return questions
}
