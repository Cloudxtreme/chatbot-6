package main

import (
    "log"
    "github.com/gocql/gocql"
)

var hostname string = "db"
var keyspace string = "question"

func DBSession() (*gocql.Session) {
    cluster := gocql.NewCluster(hostname)
	cluster.Keyspace = keyspace
    cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()
    if err != nil {
        log.Fatal(err)
    }
    return session
}
