package main

import (
	"github.com/gocql/gocql"
	"log"
)

const (
	hostname string = "db"
	keyspace string = "sentence"
)

func DBSession() *gocql.Session {
	cluster := gocql.NewCluster(hostname)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	return session
}
