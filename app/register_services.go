package app

import (
	"github.com/gin-gonic/gin"

	"github.com/DevPer/article-blog/internal/handler"
	"github.com/DevPer/article-blog/internal/logic"
)

func (a *App) RegisterServices(e *gin.Engine) {
	// register services here with concrete business logic implementations as handlers
	handler.RegisterArticleSvcHandler(e, logic.NewArticleServiceImpl(a.Config, a.RepoFactory))
}
