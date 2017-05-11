package main

import(
    "log"
    "net/http"
    "encoding/json"

    "github.com/julienschmidt/httprouter"
    "github.com/gocql/gocql"
)

var keyspace string = "questions"
var hostname string = "db"

type Question struct {
    Question string `json:question`
    Response string `json:response`
    Frequency int `json:frequency`
}

type Questions []Question

func session() (*gocql.Session) {
    cluster := gocql.NewCluster(hostname)
	cluster.Keyspace = keyspace
	session, err := cluster.CreateSession()
    defer session.Close()
    if err != nil {
        log.Fatal(err)
    }
    return session
}

func GetQuestionsEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var iter *gocql.Iter = session().Query("SELECT question, response, frequency FROM questions").Iter()

    var question string
    var response string
    var frequency int

    for iter.Scan(&question, &response, &frequency) {
        q := Question{question, response, frequency}
        json, err := json.Marshal(q)
        if err != nil {
            log.Fatal(err)
        }
        w.Write(json)
    }
}

func PostQuestionsEndpoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    w.Write([]byte("Hello " + ps.ByName("name")))
}

func main() {
    router := httprouter.New()
    router.GET("/v1/questions", GetQuestionsEndpoint)
    router.POST("/v1/questions", PostQuestionsEndpoint)

    log.Fatal(http.ListenAndServe(":8080", router))
}
