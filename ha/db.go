package ha

import (
	"database/sql"

	"github.com/bxcodec/dbresolver/v2"
)

// DB returns dbresolver DB that compatible with sql DB,
// dbresolver DB wraps connection to primaries for write operation
// and connection to replicas for read only operation
func DB(primaries []*sql.DB, replicas []*sql.DB) dbresolver.DB {
	return dbresolver.New(dbresolver.WithPrimaryDBs(primaries...),
		dbresolver.WithReplicaDBs(replicas...),
		dbresolver.WithLoadBalancer(dbresolver.RoundRobinLB))
}
