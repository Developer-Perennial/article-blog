package datasource

import (
	"database/sql"

	"github.com/DevPer/article-blog/internal/model/config/db"
)

type Sql struct {
	cfg *db.Config
	con *sql.DB
}

func NewSql(db *sql.DB, cfg *db.Config) *Sql {
	return &Sql{
		cfg: cfg,
		con: db,
	}
}

func (s *Sql) GetSqlCon() *sql.DB {
	return s.con
}

func (s *Sql) Close() error {
	return s.con.Close()
}
