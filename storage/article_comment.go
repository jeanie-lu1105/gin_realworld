package storage

import (
	"context"

	"example.com/gin_realworld/models"
)

func CreateArticleComment(ctx context.Context, articleComment *models.ArticleComment) error {
	return gormDb.WithContext(ctx).Create(articleComment).Error
}

func GetArticleCommentsByArticleId(ctx context.Context,
	articleSlug string) ([]*models.ArticleComment, error) {
	comments := make([]*models.ArticleComment, 0)
	if err := gormDb.WithContext(ctx).Select("article_comment.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article_comment.author_username = user.username").
		Joins("INNER JOIN article ON article.id = article_comment.article_id").
		Where("article.slug = ?", articleSlug).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func GetArticleCommentById(ctx context.Context, commentId int64) (*models.ArticleComment, error) {
	comment := models.ArticleComment{}
	if err := gormDb.WithContext(ctx).Select("article_comment.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article_comment.author_username = user.username").Where(
		"article.comment_id = ?", commentId).Find(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func DeleteArticleCommentById(ctx context.Context, commentId int64) error {
	return gormDb.WithContext(ctx).Where("article_comment.id = ?", commentId).Delete(&models.ArticleComment{}).Error
}
