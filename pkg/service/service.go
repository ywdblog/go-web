package service

type Service struct {
	Article ArticleService
}

func New() Service {
	return Service{
		Article: ArticleService{},
	}
}
