package datasource

import "database/sql"

type Ds interface {
	SqlDs
	Close() error
}

type SqlDs interface {
	GetSqlCon() *sql.DB
}
