package logic

import (
	"context"

	"github.com/DevPer/article-blog/internal/model/entity"
	"github.com/DevPer/article-blog/internal/repository"
)

type mockRepoFactoryImpl struct {
	mockArticleRepo
}

func (m mockRepoFactoryImpl) GetArticleRepo() repository.ArticleRepository {
	return m.mockArticleRepo
}

type mockArticleRepo struct {
	f func() (interface{}, error)
}

func (m mockArticleRepo) Create(ctx context.Context, article *entity.Article) (int64, error) {
	r, err := m.f()
	return r.(int64), err
}

func (m mockArticleRepo) GetById(ctx context.Context, i int64) (*entity.Article, error) {
	r, err := m.f()
	return r.(*entity.Article), err
}

func (m mockArticleRepo) GetAll(ctx context.Context) ([]*entity.Article, error) {
	r, err := m.f()
	return r.([]*entity.Article), err
}
