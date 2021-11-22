package stub

import (
	"context"

	"github.com/DevPer/article-blog/internal/model/errors"
)

type ArticleService interface {
	CreateArticle(ctx context.Context, request *CreateArticleRequest) (*CreateArticleResponse, *errors.Error)
	GetArticleById(ctx context.Context, request *GetArticleByIdRequest) (*GetArticleByIdResponse, *errors.Error)
	GetAllArticles(ctx context.Context, request *GetAllArticlesRequest) (*GetAllArticlesResponse, *errors.Error)
}

type CreateArticleRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Author  string `json:"author" validate:"required"`
}

type CreateArticleResponse struct {
	Id uint64 `json:"id"`
}

type ArticleResponse struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type GetArticleByIdRequest struct {
	Id int64 `json:"article_id" validate:"required" param:"article_id"`
}

type GetArticleByIdResponse struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type GetAllArticlesRequest struct{}

type GetAllArticlesResponse struct {
	Data []*ArticleResponse `json:"data"`
}
