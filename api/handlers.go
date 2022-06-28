package api

import (
	"context"

	"github.com/JesusJMM/blog-plat-go/api/posts"
	"github.com/gin-gonic/gin"
	"github.com/vingarcia/ksql"
)


func New(db *ksql.DB) gin.Engine{
  r := gin.Default()
  api := r.Group("/api")

  articleH := posts.New(db, context.Background())

  api.GET("/posts/all", articleH.AllArticles())
  api.GET("/posts/paginated", articleH.ArticlesPaginated())
  api.GET("/posts/author/:author", articleH.AllArticlesByAuthor())
  return *r
}
