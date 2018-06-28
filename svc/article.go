package svc

import (
	"fmt"
	"strconv"
	"time"

	"../api"
)

const ErrorArticleNotFound = "article with ID [%s] not found"
const ErrorDateInvalid = "date [%s] is not in the format YYYYMMDD"

type articleService struct {
	articles map[string]*api.Article
}

func ArticleService() *articleService {
	return &articleService{
		articles: make(map[string]*api.Article),
	}
}

func (a *articleService) SaveArticle(article *api.Article) (string, error) {
	if validateDate(article.Date) != nil {
		return "", fmt.Errorf(ErrorDateInvalid, article.Date)
	}

	id := strconv.Itoa(len(a.articles) + 1)
	article.ID = id
	a.articles[id] = article
	return id, nil
}

func (a *articleService) GetArticle(id string) (*api.Article, error) {
	if article, ok := a.articles[id]; ok {
		return article, nil
	}
	return nil, fmt.Errorf(ErrorArticleNotFound, id)
}

func (a *articleService) GetTag(tagName string, date string) (*api.Tag, error) {
	if validateDate(date) != nil {
		return nil, fmt.Errorf(ErrorDateInvalid, date)
	}

	articles := []string{}
	count := 0
	relatedTags := []string{}

	for id, article := range a.articles {
		if article.Date == date {
			for _, tag := range article.Tags {
				if tag == tagName {
					count++
					articles = append(articles, id)
					relatedTags = append(relatedTags, article.Tags...)
					break
				}
			}
		}
	}

	if len(articles) > 10 {
		articles = articles[len(articles)-10:]
	}

	return &api.Tag{
		Articles:    articles,
		Count:       count,
		RelatedTags: removeDuplicatesAndTag(tagName, relatedTags),
		Tag:         tagName,
	}, nil
}

func validateDate(date string) error {
	dateFmt := "20060102"
	_, err := time.Parse(dateFmt, date)

	return err
}

func removeDuplicatesAndTag(tag string, slice []string) []string {
	uniqueVals := make(map[string]bool)
	uniqueVals[tag] = true
	newSlice := []string{}
	for _, val := range slice {
		if _, dup := uniqueVals[val]; !dup {
			uniqueVals[val] = true
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}
