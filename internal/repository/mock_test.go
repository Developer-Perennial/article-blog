package repository

import (
	"context"

	"github.com/DevPer/article-blog/internal/model/entity"
)

type mockSql struct {
	f func(interface{}) (interface{}, error)
}

func (m mockSql) Insert(ctx context.Context, article *entity.Article) (int64, error) {
	r, err := m.f(article)
	return r.(int64), err
}

func (m mockSql) Get(ctx context.Context, q map[string]interface{}) ([]*entity.Article, error) {
	r, err := m.f(q)
	return r.([]*entity.Article), err
}
