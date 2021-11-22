package article

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/DevPer/article-blog/internal/model/entity"
	. "github.com/DevPer/article-blog/internal/repository/db/sql"
)

type Article interface {
	Insert(context.Context, *entity.Article) (int64, error)
	Get(context.Context, map[string]interface{}) ([]*entity.Article, error)
}

type Sql struct {
	db *sql.DB
}

func NewSqlCon(db *sql.DB) Article {
	return &Sql{
		db: db,
	}
}

func (s Sql) Insert(ctx context.Context, obj *entity.Article) (int64, error) {
	qry := `INSERT INTO article (create_at, update_at, title, content, author) VALUES (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?, ?, ?)`
	r, err := s.db.ExecContext(ctx, qry, obj.Title, obj.Content, obj.Author)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

func (s Sql) Get(ctx context.Context, searchMap map[string]interface{}) ([]*entity.Article, error) {
	qry := `SELECT id, create_at, update_at, title, content, author FROM article` + genSearchQuery(searchMap)

	r, err := s.db.QueryContext(ctx, qry)
	if err != nil {
		return nil, err
	}
	var res []*entity.Article
	for r.Next() {
		var (
			id        sql.NullInt64
			createAt  sql.NullTime
			updatedAt sql.NullTime
			title     sql.NullString
			content   sql.NullString
			author    sql.NullString
		)
		err = r.Scan(&id, &createAt, &updatedAt, &title, &content, &author)
		if err != nil {
			return nil, err
		}
		res = append(res, &entity.Article{
			Base: entity.Base{
				ID:        uint64(HandleNullInt64(id)),
				CreatedAt: HandleNullTime(createAt),
				UpdatedAt: HandleNullTime(updatedAt),
			},
			Title:   HandleNullString(title),
			Content: HandleNullString(content),
			Author:  HandleNullString(author),
		})
	}
	return res, nil
}

func genSearchQuery(searchMap map[string]interface{}) string {
	// create condition list
	var searchKeys []string

	for k, v := range searchMap {
		switch v.(type) {
		// string needs values to be wrapped in ''
		case string:
			searchKeys = append(searchKeys, fmt.Sprintf("%s='%s'", k, v))
		default:
			searchKeys = append(searchKeys, fmt.Sprintf("%s=%v", k, v))
		}
	}
	if len(searchKeys) > 0 {
		return fmt.Sprintf(" WHERE %s", strings.Join(searchKeys, " AND "))
	}
	return ""
}
