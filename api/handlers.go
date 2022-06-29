package api

import (
	"context"

	"github.com/JesusJMM/blog-plat-go/api/articles"
	"github.com/gin-gonic/gin"
	"github.com/vingarcia/ksql"
)

// Return a instance of gin.Engine with all routes
// registered
func New(db *ksql.DB) gin.Engine{
  r := gin.Default()
  api := r.Group("/api")

  articleH := articles.New(db, context.Background())

  api.GET("/articles/all", articleH.All())
  api.GET("/articles/paginated", articleH.Paginated())
  api.GET("/articles/author/:author", articleH.ByAuthorPaginated())
  return *r
}
