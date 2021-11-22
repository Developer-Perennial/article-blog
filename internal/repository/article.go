package repository

import (
	"context"
	"fmt"

	"github.com/DevPer/article-blog/internal/constants"
	"github.com/DevPer/article-blog/internal/model/datasource"
	"github.com/DevPer/article-blog/internal/model/entity"
	"github.com/DevPer/article-blog/internal/model/errors"
	"github.com/DevPer/article-blog/internal/repository/db/sql/article"
)

type ArticleRepository interface {
	Create(context.Context, *entity.Article) (int64, error)
	GetById(context.Context, int64) (*entity.Article, error)
	GetAll(ctx context.Context) ([]*entity.Article, error)
}

type ArticleRepositoryImpl struct {
	articleSqlDb article.Article
}

func NewArticleRepository(ds datasource.Ds) ArticleRepository {
	return &ArticleRepositoryImpl{
		articleSqlDb: article.NewSqlCon(ds.GetSqlCon()),
	}
}

func (a ArticleRepositoryImpl) Create(ctx context.Context, article *entity.Article) (int64, error) {
	return a.articleSqlDb.Insert(ctx, article)
}

func (a ArticleRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Article, error) {
	articles, err := a.articleSqlDb.Get(ctx, map[string]interface{}{
		constants.ArticleColumnId: id,
	})
	if err != nil {
		return nil, err
	}
	if len(articles) != 1 {
		return nil, fmt.Errorf(errors.ErrResourceNotFound)
	}
	return articles[0], nil
}

func (a ArticleRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Article, error) {
	articles, err := a.articleSqlDb.Get(ctx, nil)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
