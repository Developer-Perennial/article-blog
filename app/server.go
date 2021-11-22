package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DevPer/article-blog/config"
	"github.com/DevPer/article-blog/internal/handler/middleware"
	"github.com/DevPer/article-blog/internal/model/dto"
)

type Server interface {
	Run(context.Context) error
	ShutDown(context.Context) error
}

func GenServer(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.PanicHandler())

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.RespondSuccess(http.StatusOK, nil))
	})

	return r
}
