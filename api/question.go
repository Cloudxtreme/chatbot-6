package main

type Question struct {
    Question string `json:question`
    Response string `json:response`
    Frequency int `json:frequency`
}

type Questions []Question

type QuestionRepository struct {}

func (q *QuestionRepository) GetAll() {

}
