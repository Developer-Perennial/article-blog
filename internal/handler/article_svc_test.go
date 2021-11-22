package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/DevPer/article-blog/internal/model/dto"
	"github.com/DevPer/article-blog/internal/model/errors"
	"github.com/DevPer/article-blog/internal/model/stub"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestArticleSvcController_CreateArticle(t *testing.T) {
	tests := []struct {
		name   string
		opFunc func() (interface{}, *errors.Error)
		give   *stub.CreateArticleRequest
		want   struct {
			code int
			resp *dto.HttpResponse
		}
	}{
		{
			name: "validate failure",
			opFunc: func() (interface{}, *errors.Error) {
				return &stub.CreateArticleResponse{}, nil
			},
			give: &stub.CreateArticleRequest{},
			want: struct {
				code int
				resp *dto.HttpResponse
			}{
				code: http.StatusBadRequest,
				resp: &dto.HttpResponse{
					Status:  http.StatusBadRequest,
					Message: "required fields missing",
					Data:    nil,
				},
			},
		},
		{
			name: "logic error",
			opFunc: func() (interface{}, *errors.Error) {
				return &stub.CreateArticleResponse{}, errors.New(http.StatusInternalServerError, "test error")
			},
			give: &stub.CreateArticleRequest{
				Title:   "title",
				Content: "content",
				Author:  "author",
			},
			want: struct {
				code int
				resp *dto.HttpResponse
			}{
				code: http.StatusInternalServerError,
				resp: &dto.HttpResponse{
					Status:  http.StatusInternalServerError,
					Message: "test error",
					Data:    nil,
				},
			},
		},
		{
			name: "logic success",
			opFunc: func() (interface{}, *errors.Error) {
				return &stub.CreateArticleResponse{
					Id: 1,
				}, nil
			},
			give: &stub.CreateArticleRequest{
				Title:   "title",
				Content: "content",
				Author:  "author",
			},
			want: struct {
				code int
				resp *dto.HttpResponse
			}{
				code: http.StatusCreated,
				resp: &dto.HttpResponse{
					Status:  http.StatusCreated,
					Message: "Success",
					Data: &stub.CreateArticleResponse{
						Id: 1,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare custom handler
			router := gin.Default()

			RegisterArticleSvcHandler(router, &mockArticleSvc{
				f: tt.opFunc,
			})

			resp := httptest.NewRecorder()
			req := prepareCustomTestHttpRequest(t, tt.give, http.MethodPost, "/articles")

			router.ServeHTTP(resp, req)

			httpResp := struct {
				code int
				resp *dto.HttpResponse
			}{
				code: resp.Code,
				resp: &dto.HttpResponse{
					Status:  0,
					Message: "",
					Data:    &stub.CreateArticleResponse{},
				},
			}
			_ = json.Unmarshal(resp.Body.Bytes(), httpResp.resp)
			if !reflect.DeepEqual(tt.want, httpResp) {
				t.Errorf("Want: %#v, Got: %#v", tt.want, httpResp)
			}
		})
	}
}

func TestArticleSvcController_GetArticleById(t *testing.T) {
	tests := []struct {
		name   string
		opFunc func() (interface{}, *errors.Error)
		give   string
		want   struct {
			code int
			resp *dto.HttpResponse
		}
	}{
		{
			name: "parse failure",
			opFunc: func() (interface{}, *errors.Error) {
				return &stub.GetArticleByIdResponse{}, nil
			},
			give: "1mnb3",
			want: struct {
				code int
				resp *dto.HttpResponse
			}{
				code: http.StatusBadRequest,
				resp: &dto.HttpResponse{
					Status:  http.StatusBadRequest,
					Message: "article id must be a whole number",
					Data:    nil,
				},
			},
		},
		{
			name: "logic error",
			opFunc: func() (interface{}, *errors.Error) {
				return &stub.GetArticleByIdResponse{}, errors.New(http.StatusInternalServerError, "test error")
			},
			give: "1",
			want: struct {
				code int
				resp *dto.HttpResponse
			}{
				code: http.StatusInternalServerError,
				resp: &dto.HttpResponse{
					Status:  http.StatusInternalServerError,
					Message: "test error",
					Data:    nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare custom handler
			router := gin.Default()

			RegisterArticleSvcHandler(router, &mockArticleSvc{
				f: tt.opFunc,
			})

			resp := httptest.NewRecorder()
			req := prepareCustomTestHttpRequest(t, tt.give, http.MethodGet, "/articles/"+tt.give)

			router.ServeHTTP(resp, req)

			httpResp := struct {
				code int
				resp *dto.HttpResponse
			}{
				code: resp.Code,
				resp: &dto.HttpResponse{
					Status:  0,
					Message: "",
					Data:    []*stub.GetArticleByIdResponse{},
				},
			}
			_ = json.Unmarshal(resp.Body.Bytes(), httpResp.resp)
			if !reflect.DeepEqual(tt.want, httpResp) {
				t.Errorf("Want: %#v, Got: %#v", tt.want, httpResp)
			}
		})
	}
}

func TestArticleSvcController_GetAllArticles(t *testing.T) {
	tests := []struct {
		name   string
		opFunc func() (interface{}, *errors.Error)
		want   struct {
			code int
			resp *dto.HttpResponse
		}
	}{
		{
			name: "logic error",
			opFunc: func() (interface{}, *errors.Error) {
				return &stub.GetAllArticlesResponse{}, errors.New(http.StatusInternalServerError, "test error")
			},
			want: struct {
				code int
				resp *dto.HttpResponse
			}{
				code: http.StatusInternalServerError,
				resp: &dto.HttpResponse{
					Status:  http.StatusInternalServerError,
					Message: "test error",
					Data:    nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare custom handler
			router := gin.Default()

			RegisterArticleSvcHandler(router, &mockArticleSvc{
				f: tt.opFunc,
			})

			resp := httptest.NewRecorder()
			req := prepareCustomTestHttpRequest(t, nil, http.MethodGet, "/articles")

			router.ServeHTTP(resp, req)

			httpResp := struct {
				code int
				resp *dto.HttpResponse
			}{
				code: resp.Code,
				resp: &dto.HttpResponse{
					Status:  0,
					Message: "",
					Data:    []*stub.ArticleResponse{},
				},
			}
			t.Logf(resp.Body.String())
			_ = json.Unmarshal(resp.Body.Bytes(), httpResp.resp)
			if !reflect.DeepEqual(tt.want, httpResp) {
				t.Logf("%+v", httpResp.resp)
				t.Errorf("Want: %#v, Got: %#v", tt.want.resp, httpResp.resp)
			}
		})
	}
}

func prepareCustomTestHttpRequest(t *testing.T, data interface{}, httpMethod, url string) *http.Request {
	rawData, err := json.Marshal(&data)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	return httptest.NewRequest(httpMethod, url, bytes.NewBuffer(rawData))
}
