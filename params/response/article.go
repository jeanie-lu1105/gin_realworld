package response

import (
	"time"

	"example.com/gin_realworld/models"
)

type ListArticlesResponse struct {
	ArticlesCount int64      `json:"articlesCount"`
	Articles      []*Article `json:"articles"`
}

type Article struct {
	Author         *ArticleAuthor `json:"author"`
	Title          string         `json:"title"`
	Slug           string         `json:"slug"`
	Body           string         `json:"body"`
	Description    string         `json:"description"`
	TagList        []string       `json:"tagList"`
	Favorited      bool           `json:"favorited"`
	FavoritesCount int            `json:"favoritesCount"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

type ArticleAuthor struct {
	Bio       string `json:"bio"`
	Following bool   `json:"following"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}

func (a *Article) FromDB(dbArticle *models.Article) {
	a.Author = &ArticleAuthor{
		Bio:       dbArticle.AuthorUserBio,
		Following: false,
		Image:     dbArticle.AuthorUserImage,
		Username:  dbArticle.AuthorUsername,
	}
	a.Title = dbArticle.Title
	a.Slug = dbArticle.Slug
	a.Body = dbArticle.Body
	a.Description = dbArticle.Description
	a.TagList = dbArticle.TagList
	a.Favorited = false
	a.FavoritesCount = 0
	a.CreatedAt = dbArticle.CreatedAt
	a.UpdatedAt = dbArticle.UpdatedAt
}
