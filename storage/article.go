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

func ListArticles(ctx context.Context, req *request.ListArticleQuery) ([]*models.Article, error) {
	var articles []*models.Article
	err := listArticleTx(ctx, req).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func listArticleTx(ctx context.Context, req *request.ListArticleQuery) *gorm.DB {
	tx := gormDb.WithContext(ctx).Model(models.Article{}).
		Select("article.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article.author_username = user.username").
		Order("created_at desc")
	if req.Username != "" {
		tx = tx.Where("user.username = ?", req.Username)
	}
	if req.Tag != "" {
		// where ()
		tx = tx.Where("article.tag_list like ?", "%\""+req.Tag+"\"%")
	}
	return tx
}

func GetArticleBySlug(ctx context.Context, slug string) (*models.Article, error) {
	var article models.Article
	if err := gormDb.WithContext(ctx).Model(models.Article{}).
		Select("article.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article.author_username = user.username").
		Where("slug = ?", slug).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func GetArticleByUser(ctx context.Context, username string) (*models.Article, error) {
	var article models.Article
	if err := gormDb.WithContext(ctx).Model(models.Article{}).Select("article.*, "+
		"user.email as author_user_email, user.bio as author_user_bio, "+
		"user.image as author_user_image").Joins("left join user on article."+
		"author_username = user.username").Where("user.username = ?",
		username).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func CountArticles(ctx context.Context, req *request.ListArticleQuery) (int64, error) {
	var count int64
	if err := listArticleTx(ctx, req).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
