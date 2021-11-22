package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/DevPer/article-blog/internal/model/config/db"
	"github.com/DevPer/article-blog/internal/model/datasource"
)

func SqlConnect(cfg *db.Config) (datasource.Ds, error) {
	cfg.InjectEnvValues()
	if cfg.Validate() != nil {
		return nil, errors.New(fmt.Sprintf("DB Connect failed::missing required config"))
	}
	if cfg.GenDsn() != nil {
		return nil, errors.New(fmt.Sprintf("DB Connect failed::failed to generate DSN"))
	}
	db, err := sql.Open("mysql", cfg.Dsn)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("DB Connect failed::%s", err.Error()))
	}
	if cfg.MaxIdleCon != nil {
		db.SetMaxIdleConns(*cfg.MaxIdleCon)
	}
	if cfg.MaxOpenCon != nil {
		db.SetMaxOpenConns(*cfg.MaxOpenCon)
	}
	return datasource.NewSql(db, cfg), db.Ping()
}
