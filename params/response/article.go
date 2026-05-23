package response

import "time"

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
