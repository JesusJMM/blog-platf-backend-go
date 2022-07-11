package articles

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesusJMM/blog-plat-go/postgres/repos/articles"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/ksqltest"
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

var testArticleData = []map[string]interface{}{
	{
		"article_id":    0,
		"title": "Test Article",
		"slug":  "Article Description",
    "user_id": 3,
	},
	{
		"article_id":    1,
		"title": "Test Article",
		"slug":  "Article Description",
    "user_id": 3,
	},
}

// test the 'all' controller
func Test_AllController(t *testing.T) {
	t.Run("Should returns the entire article set", func(t *testing.T) {
		mockDB := ksql.Mock{
			QueryFn: func(ctx context.Context, record interface{}, query string, params ...interface{}) error {
        ksqltest.FillSliceWith(record, testArticleData)
				return nil
			},
		}
		r := setupTestingRouter(mockDB, context.Background(), articles.NewMockedArticleRepo())


		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/all", nil)
		r.ServeHTTP(w, req)

    var body = make(map[string][]map[string]interface{}, 3)

    err := json.Unmarshal(w.Body.Bytes(), &body)
    if err != nil {
      t.Errorf("Fail to decode the body response: %v", err)
    }
		assert.Equal(t, 200, w.Code)
    assert.NotEqual(t, len(body["articles"]), len(testArticleData))
	})
}
