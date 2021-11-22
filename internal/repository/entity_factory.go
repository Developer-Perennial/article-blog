package repository

import "github.com/DevPer/article-blog/internal/model/datasource"

type RepoFactory interface {
	GetArticleRepo() ArticleRepository
}

type RepoFactoryImpl struct {
	ds          datasource.Ds
	articleRepo ArticleRepository
}

func NewRepoFactory(ds datasource.Ds) RepoFactory {
	return &RepoFactoryImpl{
		ds:          ds,
		articleRepo: NewArticleRepository(ds),
	}
}

func (r RepoFactoryImpl) GetArticleRepo() ArticleRepository {
	return r.articleRepo
}
