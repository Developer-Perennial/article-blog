package sql

import (
	"database/sql"
	"time"
)

func HandleNullString(s sql.NullString) string {
	return s.String
}

func HandleNullInt64(s sql.NullInt64) int64 {
	return s.Int64
}

func HandleNullTime(s sql.NullTime) time.Time {
	return s.Time
}
