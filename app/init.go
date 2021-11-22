package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/DevPer/article-blog/config"
	"github.com/DevPer/article-blog/internal/model/datasource"
	"github.com/DevPer/article-blog/internal/repository"
	"github.com/DevPer/article-blog/internal/repository/db"
)

type App struct {
	s *http.Server
	*config.Config
	e *gin.Engine
	datasource.Ds
	repository.RepoFactory
}

func Init(cfgFilePath string) (*App, error) {
	a := &App{}
	var err error

	a.Config = config.LoadConfigFromFile(cfgFilePath)
	a.Ds, err = db.SqlConnect(a.Config.DbConfig)
	if err != nil {
		return nil, err
	}
	a.RepoFactory = repository.NewRepoFactory(a.Ds)

	a.e = GenServer(a.Config)

	a.RegisterServices(a.e)

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	a.ServerConfig.InjectEnvValues()
	if a.ServerConfig.Validate() != nil {
		return errors.New(fmt.Sprintf("Server start failed::missing required config"))
	}
	a.s = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", a.ServerConfig.Host, a.ServerConfig.Port),
		Handler: a.e,
	}
	fmt.Println("Starting Server::", a.s.Addr)
	err := a.s.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return errors.New(fmt.Sprintf("Server error::%s", err.Error()))
	}
	return nil
}

func (a *App) ShutDown(ctx context.Context) error {
	if a.s == nil {
		return errors.New(fmt.Sprintf("Server shutdown failed::server not initiated"))
	}
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	defer a.Ds.Close()

	err := a.s.Shutdown(ctx)
	if err != nil {
		return errors.New(fmt.Sprintf("Server shutdown error::%s", err.Error()))
	}
	return nil
}
