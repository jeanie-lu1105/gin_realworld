package response

import (
	"time"

	"example.com/gin_realworld/models"
)

type ArticleComment struct {
	Id        int64     `json:"id"`
	Author    *Author   `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (a *ArticleComment) FromDB(dbArticleComment *models.ArticleComment) {
	a.Author = &Author{
		Bio:       dbArticleComment.AuthorUserBio,
		Following: false,
		Image:     dbArticleComment.AuthorUserImage,
		Username:  dbArticleComment.AuthorUsername,
	}
	a.Id = dbArticleComment.Id
	a.Body = dbArticleComment.Body
	a.CreatedAt = dbArticleComment.CreatedAt
	a.UpdatedAt = dbArticleComment.UpdatedAt
}
