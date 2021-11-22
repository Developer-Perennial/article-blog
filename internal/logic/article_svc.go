package logic

import (
	"context"
	"strings"

	"github.com/DevPer/article-blog/config"
	"github.com/DevPer/article-blog/internal/model/entity"
	"github.com/DevPer/article-blog/internal/model/errors"
	"github.com/DevPer/article-blog/internal/model/stub"
	"github.com/DevPer/article-blog/internal/repository"
)

type ArticleServiceImpl struct {
	*config.Config
	repository.RepoFactory
}

func NewArticleServiceImpl(cfg *config.Config, repoFactory repository.RepoFactory) *ArticleServiceImpl {
	return &ArticleServiceImpl{
		Config:      cfg,
		RepoFactory: repoFactory,
	}
}

func (a ArticleServiceImpl) CreateArticle(ctx context.Context, request *stub.CreateArticleRequest) (*stub.CreateArticleResponse, *errors.Error) {
	if strings.TrimSpace(request.Title) == "" ||
		strings.TrimSpace(request.Author) == "" ||
		strings.TrimSpace(request.Content) == "" {
		return nil, errors.NewBadRequest400("fields cannot be empty")
	}

	article := entity.NewArticle(request.Title, request.Content, request.Author)
	id, err := a.GetArticleRepo().Create(ctx, article)
	if err != nil {
		return nil, errors.NewInternalServerError500(err.Error())
	}
	return &stub.CreateArticleResponse{
		Id: uint64(id),
	}, nil
}

func (a ArticleServiceImpl) GetArticleById(ctx context.Context, request *stub.GetArticleByIdRequest) (*stub.GetArticleByIdResponse, *errors.Error) {
	article, err := a.GetArticleRepo().GetById(ctx, request.Id)
	if err != nil {
		if err.Error() == errors.ErrResourceNotFound {
			return nil, errors.NewResourceNotFound404("article not found")
		}
		return nil, errors.NewInternalServerError500(err.Error())
	}
	return &stub.GetArticleByIdResponse{
		Id:      int64(article.ID),
		Title:   article.Title,
		Content: article.Content,
		Author:  article.Author,
	}, nil
}

func (a ArticleServiceImpl) GetAllArticles(ctx context.Context, request *stub.GetAllArticlesRequest) (*stub.GetAllArticlesResponse, *errors.Error) {
	articles, err := a.GetArticleRepo().GetAll(ctx)
	if err != nil {
		return nil, errors.NewInternalServerError500(err.Error())
	}

	resp := &stub.GetAllArticlesResponse{}

	for _, article := range articles {
		resp.Data = append(resp.Data, &stub.ArticleResponse{
			Id:      int64(article.ID),
			Title:   article.Title,
			Content: article.Content,
			Author:  article.Author,
		})
	}
	return resp, nil
}
