package request

type ListArticleQuery struct {
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
	Tag    string `form:"tag"`
}

type CreateArticleRequest struct {
	Article Article `json:"article"`
}

type Article struct {
	Title       string   `json:"title"`
	Body        string   `json:"body"`
	Description string   `json:"description"`
	TagList     []string `json:"tagList"`
}
