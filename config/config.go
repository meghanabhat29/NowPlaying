package config

import (
	"log"

	"github.com/gocql/gocql"
)

//KeyspaceConnection ...
type KeyspaceConnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var connection KeyspaceConnection

//Setupconnection ...
func Setupconnection() {
	var err error
	connection.cluster = gocql.NewCluster("127.0.0.1")
	connection.cluster.Consistency = gocql.Quorum
	connection.cluster.Keyspace = "playlist"
	connection.session, err = connection.cluster.CreateSession()
	if err != nil {
		log.Println(err)
	}
	//defer connection.session.Close()
	//return
}
