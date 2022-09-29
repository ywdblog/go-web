package dao

import (
	"context"
)

type Article struct {
}

func NewArticle(ctx context.Context) (*Article, error) {

	article := &Article{}
	return article, nil
}
