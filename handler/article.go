package handler

import (
	"net/http"
	"strings"

	"example.com/gin_realworld/logger"
	"example.com/gin_realworld/models"
	"example.com/gin_realworld/params/request"
	"example.com/gin_realworld/params/response"
	"example.com/gin_realworld/security"
	"example.com/gin_realworld/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddArticlesHandler(r *gin.Engine) {
	articlesGroup := r.Group("/api/articles")
	articlesGroup.GET("", listArticles)
	articlesGroup.GET("/:slug", getArticle)
	//articlesGroup.Use(middlewares.AuthMiddleware)
	articlesGroup.POST("", createArticles)
	articlesGroup.PUT("/:slug", editArticles)
	articlesGroup.DELETE("/:slug", deleteArticles)
}

func listArticles(ctx *gin.Context) {
	log := logger.New(ctx)
	var req request.ListArticleQuery
	if err := ctx.Bind(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if req.Limit == 0 {
		req.Limit = 10
		req.Offset = 0
	}
	articles, err := storage.ListArticles(ctx, &req)
	if err != nil {
		log.WithError(err).Infof("list articles failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	total, err := storage.CountArticles(ctx)
	if err != nil {
		log.WithError(err).Infof("count articles failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var resp response.ListArticlesResponse
	resp.ArticlesCount = total
	for _, article := range articles {
		resp.Articles = append(resp.Articles, &response.Article{
			Author: &response.ArticleAuthor{
				Bio:       article.AuthorUserBio,
				Following: false,
				Image:     article.AuthorUserImage,
				Username:  article.AuthorUsername,
			},
			Title:          article.Title,
			Slug:           article.Slug,
			Body:           article.Body,
			Description:    article.Description,
			TagList:        article.TagList,
			Favorited:      false,
			FavoritesCount: 0,
			CreatedAt:      article.CreatedAt,
			UpdatedAt:      article.UpdatedAt,
		})
	}
	ctx.JSON(http.StatusOK, resp)
}

func createArticles(ctx *gin.Context) {
	var req request.CreateArticleRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	slug := strings.ReplaceAll(req.Article.Title, " ", "-") + "-" + uuid.NewString()
	if err := storage.CreateArticle(ctx, &models.Article{
		AuthorUsername: security.GetCurrentUsername(ctx),
		Title:          req.Article.Title,
		Slug:           slug,
		Body:           req.Article.Body,
		Description:    req.Article.Description,
		TagList:        req.Article.TagList,
	}); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusCreated, map[string]interface{}{"article": respArticle})
}

func getArticle(ctx *gin.Context) {
	slug := ctx.Param("slug")
	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusOK, map[string]interface{}{"article": respArticle})
}

func editArticles(ctx *gin.Context) {
	oldSlug := ctx.Param("slug")
	var req request.CreateArticleRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	slug := strings.ReplaceAll(req.Article.Title, " ", "-") + "-" + uuid.NewString()
	if err := storage.UpdateArticle(ctx, oldSlug, &models.Article{
		AuthorUsername: security.GetCurrentUsername(ctx),
		Title:          req.Article.Title,
		Slug:           slug,
		Body:           req.Article.Body,
		Description:    req.Article.Description,
		TagList:        req.Article.TagList,
	}); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusOK, map[string]interface{}{"article": respArticle})
}

func deleteArticles(ctx *gin.Context) {
	slug := ctx.Param("slug")
	oldArticle, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if oldArticle.AuthorUsername != security.GetCurrentUsername(ctx) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	err = storage.DeleteArticle(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
