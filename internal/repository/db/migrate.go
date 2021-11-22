package db

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/DevPer/article-blog/internal/model/config/db"
	"github.com/DevPer/article-blog/internal/model/datasource"
	"github.com/DevPer/article-blog/internal/model/entity"
)

func AutoMigrate(cfg *db.Config, ds datasource.Ds) error {
	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: ds.GetSqlCon(),
	}), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		Logger: logger.Default.LogMode(func() logger.LogLevel {
			if cfg.Debug {
				return logger.Info
			}
			return logger.Warn
		}()),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("DB migrate failed::GORM connect Failure::%s", err.Error()))
	}
	err = gormDb.AutoMigrate(entity.Entities...)
	if err != nil {
		return errors.New(fmt.Sprintf("DB migrate failed::%s", err.Error()))
	}
	return nil
}
