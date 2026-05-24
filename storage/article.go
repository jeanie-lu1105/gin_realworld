package storage

import (
	"context"

	"example.com/gin_realworld/models"
	"example.com/gin_realworld/params/request"
	"gorm.io/gorm"
)

func CreateArticle(ctx context.Context, article *models.Article) error {
	return gormDb.WithContext(ctx).Create(article).Error
}

func UpdateArticle(ctx context.Context, oldSlug string, article *models.Article) error {
	return gormDb.WithContext(ctx).Model(article).Where("slug = ?", oldSlug).Updates(article).Error
}

func DeleteArticle(ctx context.Context, slug string) error {
	return gormDb.WithContext(ctx).Where("slug = ?", slug).Delete(&models.Article{}).Error
}

func CountArticles(ctx context.Context) (int64, error) {
	var count int64
	if err := gormDb.WithContext(ctx).Model(models.Article{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func ListArticles(ctx context.Context, req *request.ListArticleQuery) ([]*models.Article, error) {
	var articles []*models.Article
	err := ListArticleTx(ctx, req).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func ListArticleTx(ctx context.Context, req *request.ListArticleQuery) *gorm.DB {
	tx := gormDb.WithContext(ctx).Model(models.Article{}).Select("article.*,  " +
		"user.email as author_user_email, user.bio as author_user_bio, " +
		"user.image as author_user_image").Joins("left join user on article." +
		"author_username = user.username").Order("article.created_at desc").Offset(req.
		Offset).Limit(req.Limit)
	if req.Tag != "" {
		tx = tx.Where("article.tag_list like ?", "%\""+req.Tag+"\"%")
	}
	return tx
}

func GetArticleBySlug(ctx context.Context, slug string) (*models.Article, error) {
	var article models.Article
	if err := gormDb.WithContext(ctx).Model(models.Article{}).Select("article.*, "+
		"user.email as author_user_email, user.bio as author_user_bio, "+
		"user.image as author_user_image").Joins("left join user on article."+
		"author_username = user.username").Where("article.slug = ?",
		slug).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
