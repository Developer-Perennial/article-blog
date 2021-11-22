package handler

import (
	"context"

	"github.com/DevPer/article-blog/internal/model/errors"
	"github.com/DevPer/article-blog/internal/model/stub"
)

type mockArticleSvc struct {
	f func() (interface{}, *errors.Error)
}

func (m mockArticleSvc) CreateArticle(ctx context.Context, request *stub.CreateArticleRequest) (*stub.CreateArticleResponse, *errors.Error) {
	r, err := m.f()
	return r.(*stub.CreateArticleResponse), err
}

func (m mockArticleSvc) GetArticleById(ctx context.Context, request *stub.GetArticleByIdRequest) (*stub.GetArticleByIdResponse, *errors.Error) {
	r, err := m.f()
	return r.(*stub.GetArticleByIdResponse), err
}

func (m mockArticleSvc) GetAllArticles(ctx context.Context, request *stub.GetAllArticlesRequest) (*stub.GetAllArticlesResponse, *errors.Error) {
	r, err := m.f()
	return r.(*stub.GetAllArticlesResponse), err
}


