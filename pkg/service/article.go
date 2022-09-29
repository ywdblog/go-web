package service

import (
	"context"

	. "github.com/xiwujie/article/pkg/entity"
)

type ArticleSyncJobRequest struct {
	BaseRequest

	Limit   int    `json:"limit"`
	JobName string `json:"job_name"`
}

type ArticleSyncJobResponse struct {
	BaseResponse
}

type ArticleService struct {
}

func (s *ArticleService) MGet(ctx context.Context, req ArticleInfoMGetRequest) (interface{}, error) {
	var resp ArticleInfoMGetResponse

	if req.ArticleIds == nil || len(req.ArticleIds) < 1 {
		return resp, nil
	}

	return resp, nil
}
