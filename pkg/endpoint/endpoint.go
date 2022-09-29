package endpoint

import "github.com/xiwujie/article/pkg/service"

type Endpoints struct {
	Article ArticleEndpoints
}

// New ...
func New(s service.Service) Endpoints {

	return Endpoints{
		Article: NewArticleEndpoints(s),
	}
}
