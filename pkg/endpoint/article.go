package endpoint

import (
	"context"

	kitendpoint "github.com/go-kit/kit/endpoint"
	. "github.com/xiwujie/article/pkg/entity"
	"github.com/xiwujie/article/pkg/service"
)

type ArticleEndpoints struct {
	MGet kitendpoint.Endpoint
}

func NewArticleEndpoints(s service.Service) ArticleEndpoints {
	return ArticleEndpoints{
		MGet: MakeArticleMGetEndpoint(s),
	}
}

func MakeArticleMGetEndpoint(s service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArticleInfoMGetRequest)
		return s.Article.MGet(ctx, req)
	}
}
