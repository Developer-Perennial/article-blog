package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DevPer/article-blog/internal/model/dto"
	"github.com/DevPer/article-blog/internal/model/errors"
)

func PanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				customErr := errors.NewInternalServerError500(fmt.Errorf("%+v", p).Error())
				if e, ok := p.(error); ok {
					customErr = errors.NewInternalServerError500(e.Error())
				}
				c.JSON(http.StatusInternalServerError, dto.RespondError(customErr.Code, customErr.Msg))
			}
		}()
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}
