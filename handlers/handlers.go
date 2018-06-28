package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../api"
)

type GetArticleService interface {
	GetArticle(id string) (*api.Article, error)
}

type PostArticleService interface {
	SaveArticle(article *api.Article) (string, error)
}

type GetTagService interface {
	GetTag(tagName string, date string) (*api.Tag, error)
}

func postArticleHandler(articlePoster PostArticleService) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		var article api.Article
		err = json.Unmarshal(body, &article)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		id, err := articlePoster.SaveArticle(&article)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(fmt.Sprintf("Article with ID [%s] created", id)))
	})
}

func getArticleHandler(articleGetter GetArticleService) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		id := req.URL.Query().Get(":id")
		article, err := articleGetter.GetArticle(id)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte(err.Error()))
			return
		}

		articleJson, err := json.Marshal(article)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(articleJson))
	})
}

func getTagHandler(tagGetter GetTagService) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		tagName := req.URL.Query().Get(":tagName")
		date := req.URL.Query().Get(":date")

		tag, err := tagGetter.GetTag(tagName, date)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		tagJson, err := json.Marshal(tag)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(tagJson))
	})
}
