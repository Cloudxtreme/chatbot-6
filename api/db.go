package main

import (
    "github.com/gocql/gocql"
)

var keyspace string = "questions"
var hostname string = "db"

func DBSession() (*gocql.Session) {
    cluster := gocql.NewCluster(hostname)
	cluster.Keyspace = keyspace
	session, err := cluster.CreateSession()
    defer session.Close()
    if err != nil {
        log.Fatal(err)
    }
    return session
}
