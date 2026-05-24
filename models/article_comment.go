package models

import "time"

type ArticleComment struct {
	Id             int64  `gorm:"primaryKey"`
	AuthorUsername string `gorm:"column:author_username"`
	Body           string
	ArticleId      int64
	CreatedAt      time.Time
	UpdatedAt      time.Time

	AuthorUserEmail string `gorm:"->"`
	AuthorUserImage string `gorm:"->"`
	AuthorUserBio   string `gorm:"->"`
}

func (ac ArticleComment) TableName() string {
	return "article_comment"
}
