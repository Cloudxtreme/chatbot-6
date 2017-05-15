package main

import(
    "log"
    "net/http"

    "github.com/julienschmidt/httprouter"
)

func GetIndexEndPoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    w.Write([]byte("Hello!"))
}

func main() {
    router := httprouter.New()
    router.GET("/", GetIndexEndPoint)

    log.Fatal(http.ListenAndServe(":8080", router))
}
