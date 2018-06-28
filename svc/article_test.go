package svc_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"../api"
	"../svc"
)

func TestArticleService_GetArticle(t *testing.T) {
	var articleSvc = svc.ArticleService()

	var mockArticle = &api.Article{
		ID:   "1",
		Date: "20080808",
	}

	articleSvc.SaveArticle(mockArticle)

	var testCases = []struct {
		Name    string
		Article *api.Article
		ID      string
		Err     error
	}{
		{
			Name:    "Successful",
			Article: mockArticle,
			ID:      "1",
		},
		{
			Name: "Not Found",
			ID:   "2",
			Err:  errors.New("article with ID [2] not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			article, err := articleSvc.GetArticle(tc.ID)
			assert.Equal(t, tc.Article, article)

			if tc.Err == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tc.Err, err)
			}
		})
	}
}

func TestArticleService_PostArticle(t *testing.T) {
	var testCases = []struct {
		Name    string
		Article *api.Article
		ID      string
		Err     error
	}{
		{
			Name: "Successful",
			Article: &api.Article{
				Date: "20080808",
			},
			ID: "1",
		},
		{
			Name: "Date in invalid format",
			Article: &api.Article{
				ID:   "1",
				Date: "2008-08-08",
			},
			Err: errors.New("date [2008-08-08] is not in the format YYYYMMDD"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var articleSvc = svc.ArticleService()
			id, err := articleSvc.SaveArticle(tc.Article)
			assert.Equal(t, tc.ID, id)

			if tc.Err == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tc.Err, err)
			}
		})
	}
}
