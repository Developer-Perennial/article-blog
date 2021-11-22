package repository

import (
	"context"
	"fmt"
	"github.com/DevPer/article-blog/internal/constants"
	"github.com/DevPer/article-blog/internal/model/entity"
	"github.com/DevPer/article-blog/internal/model/errors"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestArticleRepositoryImpl_Create(t *testing.T) {
	testRepoFac := &RepoFactoryImpl{
		ds: nil,
		articleRepo: &ArticleRepositoryImpl{
			articleSqlDb: mockSql{
				f: func(_ interface{}) (interface{}, error) {
					return int64(1), nil
				},
			},
		},
	}
	resp, err := testRepoFac.GetArticleRepo().Create(context.Background(), entity.NewArticle("title", "content", "author"))
	if err != nil || resp != 1 {
		t.Errorf("Unexpected error")
	}
}

func TestArticleRepositoryImpl_GetById(t *testing.T) {
	testRepoFac := &RepoFactoryImpl{
		ds: nil,
		articleRepo: &ArticleRepositoryImpl{
			articleSqlDb: mockSql{
				f: func(i interface{}) (interface{}, error) {
					switch i.(map[string]interface{})[constants.ArticleColumnId].(int64) {
					case 0:
						return []*entity.Article{}, fmt.Errorf("test error")
					case 1:
						return []*entity.Article{
							{
								Base:    entity.Base{},
								Title:   "",
								Content: "",
								Author:  "",
							},
							{
								Base:    entity.Base{},
								Title:   "",
								Content: "",
								Author:  "",
							},
						}, nil
					case 2:
						return []*entity.Article{
							{
								Base:    entity.Base{},
								Title:   "",
								Content: "",
								Author:  "",
							},
						}, nil
					}
					return []*entity.Article{}, nil
				},
			},
		},
	}

	tests := []struct {
		name  string
		give  int64
		check func(*testing.T, *entity.Article, error)
	}{
		{
			name: "error",
			give: 0,
			check: func(t *testing.T, article *entity.Article, err error) {
				if !reflect.DeepEqual(err, fmt.Errorf("test error")) {
					t.Errorf("Unexpected error")
				}
			},
		},
		{
			name: "not found",
			give: 1,
			check: func(t *testing.T, article *entity.Article, err error) {
				if !reflect.DeepEqual(err, fmt.Errorf(errors.ErrResourceNotFound)) {
					t.Errorf("Unexpected error")
				}
			},
		},
		{
			name: "success",
			give: 2,
			check: func(t *testing.T, article *entity.Article, err error) {
				if article == nil || err != nil {
					t.Errorf("Unexpected error")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testRepoFac.GetArticleRepo().GetById(context.Background(), tt.give)
			tt.check(t, resp, err)
		})
	}
}

func TestArticleRepositoryImpl_GetAll(t *testing.T) {
	getErr := false
	testRepoFac := &RepoFactoryImpl{
		ds: nil,
		articleRepo: &ArticleRepositoryImpl{
			articleSqlDb: mockSql{
				f: func(_ interface{}) (interface{}, error) {
					if getErr {
						return []*entity.Article{}, fmt.Errorf("test error")
					}
					return []*entity.Article{
						{
							Base:    entity.Base{},
							Title:   "",
							Content: "",
							Author:  "",
						},
					}, nil
				},
			},
		},
	}
	tests := []struct {
		name      string
		check     func(*testing.T, []*entity.Article, error)
		setupFunc func()
	}{
		{
			name: "success",
			check: func(t *testing.T, articles []*entity.Article, err error) {
				if len(articles) <= 0 || err != nil {
					t.Errorf("Unexpected error")
				}
			},
			setupFunc: func() {},
		},
		{
			name: "error",
			check: func(t *testing.T, articles []*entity.Article, err error) {
				if !reflect.DeepEqual(err, fmt.Errorf("test error")) {
					t.Errorf("Unexpected error")
				}
			},
			setupFunc: func() {
				getErr = true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFunc != nil {
				tt.setupFunc()
			}
			resp, err := testRepoFac.GetArticleRepo().GetAll(context.Background())
			tt.check(t, resp, err)
		})
	}
}
