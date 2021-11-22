package article

import (
	"context"
	"os"
	"testing"

	"github.com/DevPer/article-blog/internal/constants"
	dbCfg "github.com/DevPer/article-blog/internal/model/config/db"
	"github.com/DevPer/article-blog/internal/model/entity"
	"github.com/DevPer/article-blog/internal/repository/db"
)

var testSql Article

func TestMain(m *testing.M) {
	ds, _ := db.SqlConnect(&dbCfg.Config{
		Username: "root",
		Password: "123456",
		Host:     "0.0.0.0",
		Port:     "3306",
		Database: "blog_system",
	})
	testSql = NewSqlCon(ds.GetSqlCon())
	testRunCode := m.Run()
	ds.Close()
	os.Exit(testRunCode)
}

func TestSql_Insert(t *testing.T) {
	id, err := testSql.Insert(context.Background(), entity.NewArticle("title", "content", "author"))
	if id <= 0 || err != nil {
		t.Errorf("Unexpected error for happy path")
	}
}

func TestSql_Get(t *testing.T) {
	id, err := testSql.Insert(context.Background(), entity.NewArticle("title", "content", "author"))
	if id <= 0 || err != nil {
		t.Errorf("Unexpected error for happy path")
	}
	articles, err := testSql.Get(context.Background(), map[string]interface{}{
		constants.ArticleColumnId:    id,
		constants.ArticleColumnTitle: "title",
	})
	if err != nil || len(articles) != 1 {
		t.Errorf("Unexpected error for happy path")
	}
}
