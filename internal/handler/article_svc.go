package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/DevPer/article-blog/internal/model/dto"
	"github.com/DevPer/article-blog/internal/model/stub"
	"github.com/DevPer/article-blog/internal/util"
)

type ArticleSvcController struct {
	svc stub.ArticleService
}

func RegisterArticleSvcHandler(e *gin.Engine, svc stub.ArticleService) {
	// create article service controller instance
	ctrl := ArticleSvcController{svc}

	// bind handler routes for article service
	e.POST("/articles", ctrl.CreateArticle)
	e.GET("/articles/:article_id", ctrl.GetArticleById)
	e.GET("/articles", ctrl.GetAllArticles)

}

func (ctrl *ArticleSvcController) CreateArticle(c *gin.Context) {
	req := &stub.CreateArticleRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, dto.RespondError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := util.ValidateRequiredFields(req); err != nil {
		c.JSON(http.StatusBadRequest, dto.RespondError(http.StatusBadRequest, "required fields missing"))
		return
	}
	resp, err := ctrl.svc.CreateArticle(c.Request.Context(), req)
	if err != nil {
		c.JSON(err.Code, dto.RespondError(err.Code, err.Msg))
		return
	}
	c.JSON(http.StatusCreated, dto.RespondSuccess(http.StatusCreated, resp))
}

func (ctrl *ArticleSvcController) GetArticleById(c *gin.Context) {
	req := &stub.GetArticleByIdRequest{}
	if id, err := strconv.ParseInt(c.Param("article_id"), 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, dto.RespondError(http.StatusBadRequest, "article id must be a whole number"))
		return
	} else {
		req.Id = id
	}
	if err := util.ValidateRequiredFields(req); err != nil {
		c.JSON(http.StatusBadRequest, dto.RespondError(http.StatusBadRequest, "required fields missing"))
		return
	}
	resp, err := ctrl.svc.GetArticleById(c.Request.Context(), req)
	if err != nil {
		c.JSON(err.Code, dto.RespondError(err.Code, err.Msg))
		return
	}
	c.JSON(http.StatusOK, dto.RespondSuccess(http.StatusOK, []interface{}{resp}))
}

func (ctrl *ArticleSvcController) GetAllArticles(c *gin.Context) {
	req := &stub.GetAllArticlesRequest{}
	resp, err := ctrl.svc.GetAllArticles(c.Request.Context(), req)
	if err != nil {
		c.JSON(err.Code, dto.RespondError(err.Code, err.Msg))
		return
	}
	c.JSON(http.StatusOK, dto.RespondSuccess(http.StatusOK, resp.Data))
}
