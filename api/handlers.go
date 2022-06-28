package api

import (
	"context"

	"github.com/JesusJMM/blog-plat-go/api/posts"
	"github.com/gin-gonic/gin"
	"github.com/vingarcia/ksql"
)

// Return a instance of gin.Engine with all routes
// registered
func New(db *ksql.DB) gin.Engine{
  r := gin.Default()
  api := r.Group("/api")

  articleH := posts.New(db, context.Background())

  api.GET("/posts/all", articleH.All())
  api.GET("/posts/paginated", articleH.Paginated())
  api.GET("/posts/author/:author", articleH.ByAuthorPaginated())
  return *r
}
