package main

import(
    "log"
    "net/http"
    "encoding/json"

    "github.com/julienschmidt/httprouter"
)

var q *QuestionRepository = new(QuestionRepository)

func GetIndexEndPoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    w.Write([]byte("Hello!"))
}

func GetQuestionEndPoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    questions := q.GetAll(ps.ByName("question"))
    encoder := json.NewEncoder(w)
    encoder.Encode(questions)
}

func main() {
    router := httprouter.New()
    router.GET("/", GetIndexEndPoint)
    router.GET("/questions", GetQuestionEndPoint)

    log.Fatal(http.ListenAndServe(":8080", router))
}
