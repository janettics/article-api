package handlers

import "github.com/bmizerany/pat"

type ArticleService interface {
	GetArticleService
	PostArticleService
	GetTagService
}

type RouterParams struct {
	ArticleService ArticleService
}

func SetupRouter(params RouterParams) *pat.PatternServeMux {
	r := pat.New()

	r.Post("/articles", postArticleHandler(params.ArticleService))
	r.Get("/articles/:id", getArticleHandler(params.ArticleService))
	r.Get("/tags/:tagName/:date", getTagHandler(params.ArticleService))

	return r
}
