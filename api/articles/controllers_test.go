package articles

import (
	"context"
	// "net/http"
	// "net/http/httptest"
	// "testing"

	"github.com/JesusJMM/blog-plat-go/postgres/repos/articles"
	"github.com/gin-gonic/gin"
	// "github.com/stretchr/testify/assert"
	"github.com/vingarcia/ksql"
)

func setupTestingRouter(db ksql.Provider, ctx context.Context, articleRepo articles.ArticleRepository) *gin.Engine {
  controllers := New(db, ctx, articleRepo)
  r := gin.Default()
  r.GET("/all", controllers.All())
  r.GET("/paginated", controllers.Paginated())
  r.GET("/byAuthorPaginated", controllers.ByAuthorPaginated())
  r.GET("/one", controllers.OneArticle())
  return r
}
