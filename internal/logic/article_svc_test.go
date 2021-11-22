package logic

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/DevPer/article-blog/internal/model/entity"
	"github.com/DevPer/article-blog/internal/model/errors"
	"github.com/DevPer/article-blog/internal/model/stub"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestArticleServiceImpl_CreateArticle(t *testing.T) {
	tests := []struct {
		name   string
		give   *stub.CreateArticleRequest
		check  func(*testing.T, *stub.CreateArticleResponse, *errors.Error)
		opFunc func() (interface{}, error)
	}{
		{
			name: "happy path",
			give: &stub.CreateArticleRequest{
				Title:   "title",
				Content: "content",
				Author:  "author",
			},
			check: func(t *testing.T, r *stub.CreateArticleResponse, err *errors.Error) {
				if !reflect.DeepEqual(r, &stub.CreateArticleResponse{Id: 1}) &&
					!reflect.DeepEqual(err, nil) {
					t.Errorf("Unexpected error::%+v", err)
				}
			},
			opFunc: func() (interface{}, error) {
				return int64(1), nil
			},
		},
		{
			name: "empty field",
			give: &stub.CreateArticleRequest{
				Title:   " ",
				Content: "content",
				Author:  "author",
			},
			check: func(t *testing.T, r *stub.CreateArticleResponse, err *errors.Error) {
				if !reflect.DeepEqual(err, errors.NewBadRequest400("fields cannot be empty")) {
					t.Errorf("Unexpected error::%+v", err)
				}
			},
			opFunc: func() (interface{}, error) {
				return nil, nil
			},
		},
		{
			name: "internal server error",
			give: &stub.CreateArticleRequest{
				Title:   "title",
				Content: "content",
				Author:  "author",
			},
			check: func(t *testing.T, r *stub.CreateArticleResponse, err *errors.Error) {
				if !reflect.DeepEqual(err, errors.NewInternalServerError500("test error")) {
					t.Errorf("Unexpected error::%+v", err)
				}
			},
			opFunc: func() (interface{}, error) {
				return int64(1), fmt.Errorf("test error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleSvc := NewArticleServiceImpl(nil, &mockRepoFactoryImpl{
				mockArticleRepo{
					f: tt.opFunc,
				},
			})
			r, err := articleSvc.CreateArticle(context.Background(), tt.give)
			tt.check(t, r, err)
		})
	}
}

func TestMockArticleRepo_GetById(t *testing.T) {
	tests := []struct {
		name   string
		give   *stub.GetArticleByIdRequest
		check  func(*testing.T, *stub.GetArticleByIdResponse, *errors.Error)
		opFunc func() (interface{}, error)
	}{
		{
			name: "happy path",
			give: &stub.GetArticleByIdRequest{
				Id: 1,
			},
			check: func(t *testing.T, r *stub.GetArticleByIdResponse, err *errors.Error) {
				if !reflect.DeepEqual(r, &stub.GetArticleByIdResponse{
					Id:      1,
					Title:   "title",
					Content: "content",
					Author:  "author",
				}) &&
					!reflect.DeepEqual(err, nil) {
					t.Errorf("Unexpected error::%+v", err)
				}
			},
			opFunc: func() (interface{}, error) {
				return &entity.Article{
					Base: entity.Base{
						ID: 1,
					},
					Title:   "title",
					Content: "content",
					Author:  "author",
				}, nil
			},
		},
		{
			name: "internal server error",
			give: &stub.GetArticleByIdRequest{
				Id: 1,
			},
			check: func(t *testing.T, r *stub.GetArticleByIdResponse, err *errors.Error) {
				if !reflect.DeepEqual(err, errors.NewInternalServerError500("test error")) {
					t.Errorf("Unexpected error::%+v", err)
				}
			},
			opFunc: func() (interface{}, error) {
				return &entity.Article{}, fmt.Errorf("test error")
			},
		},
		{
			name: "article not found",
			give: &stub.GetArticleByIdRequest{
				Id: 1,
			},
			check: func(t *testing.T, r *stub.GetArticleByIdResponse, err *errors.Error) {
				if !reflect.DeepEqual(err, errors.NewResourceNotFound404("article not found")) {
					t.Errorf("Unexpected error::%+v", err)
				}
			},
			opFunc: func() (interface{}, error) {
				return &entity.Article{}, fmt.Errorf(errors.ErrResourceNotFound)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleSvc := NewArticleServiceImpl(nil, &mockRepoFactoryImpl{
				mockArticleRepo{
					f: tt.opFunc,
				},
			})
			r, err := articleSvc.GetArticleById(context.Background(), tt.give)
			tt.check(t, r, err)
		})
	}
}

func TestArticleServiceImpl_GetAllArticles(t *testing.T) {
	tests := []struct {
		name   string
		give   *stub.GetAllArticlesRequest
		check  func(*testing.T, *stub.GetAllArticlesResponse, *errors.Error)
		opFunc func() (interface{}, error)
	}{
		{
			name: "happy path",
			give: &stub.GetAllArticlesRequest{},
			check: func(t *testing.T, r *stub.GetAllArticlesResponse, err *errors.Error) {
				if !reflect.DeepEqual(r, &stub.GetAllArticlesResponse{
					Data: []*stub.ArticleResponse{
						{
							Id:      1,
							Title:   "title",
							Content: "content",
							Author:  "author",
						},
						{
							Id:      2,
							Title:   "title",
							Content: "content",
							Author:  "author",
						},
					},
				}) &&
					!reflect.DeepEqual(err, nil) {
					t.Errorf("Unexpected error::%+v", err)
				}
			},
			opFunc: func() (interface{}, error) {
				return []*entity.Article{
					{
						Base: entity.Base{
							ID: 1,
						},
						Title:   "title",
						Content: "content",
						Author:  "author",
					},
					{
						Base: entity.Base{
							ID: 2,
						},
						Title:   "title",
						Content: "content",
						Author:  "author",
					},
				}, nil
			},
		},
		{
			name: "internal server error",
			give: &stub.GetAllArticlesRequest{},
			check: func(t *testing.T, r *stub.GetAllArticlesResponse, err *errors.Error) {
				if !reflect.DeepEqual(err, errors.NewInternalServerError500("test error")) {
					t.Errorf("Unexpected error::%+v", err)
				}
			},
			opFunc: func() (interface{}, error) {
				return []*entity.Article{}, fmt.Errorf("test error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleSvc := NewArticleServiceImpl(nil, &mockRepoFactoryImpl{
				mockArticleRepo{
					f: tt.opFunc,
				},
			})
			r, err := articleSvc.GetAllArticles(context.Background(), tt.give)
			tt.check(t, r, err)
		})
	}
}
