package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// not instance but
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	//s, e := cluster.CreateSession()
	////defer s.Close()
	//if e != nil {
	//	fmt.Printf("\n\n\n\nEXTRATERRITORIALITY")
	//	panic(e)
	//}
	//fmt.Println(s, "ok")

}

func GetSesion() (*gocql.Session, error) {
	return cluster.CreateSession()
}
