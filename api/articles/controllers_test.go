package articles

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesusJMM/blog-plat-go/postgres/repos/articles"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/ksqltest"
)

func setupTestingRouter(db ksql.Provider) *gin.Engine {
	controllers := New(db, context.Background(), articles.NewMockedArticleRepo())
	r := gin.Default()
	r.GET("/all", controllers.All())
	r.GET("/paginated", controllers.Paginated())
	r.GET("/byAuthorPaginated", controllers.ByAuthorPaginated())
	r.GET("/one", controllers.OneArticle())
	return r
}

var testArticleData = []map[string]interface{}{
	{
		"article_id": 0,
		"title":      "Test Article",
		"slug":       "Article Description",
		"user_id":    3,
	},
	{
		"article_id": 1,
		"title":      "Test Article",
		"slug":       "Article Description",
		"user_id":    3,
	},
}

type commonBodyType map[string][]map[string]interface{}

// test the 'all' controller
func Test_AllController(t *testing.T) {
	t.Run("Should returns the entire article set", func(t *testing.T) {
		mockDB := ksql.Mock{
			QueryFn: func(ctx context.Context, record interface{}, query string, params ...interface{}) error {
				ksqltest.FillSliceWith(record, testArticleData)
				return nil
			},
		}
		r := setupTestingRouter(mockDB)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/all", nil)
		r.ServeHTTP(w, req)

		var body = make(commonBodyType, 3)

		err := json.Unmarshal(w.Body.Bytes(), &body)
		if err != nil {
			t.Fatalf("Fail to decode the body response: %v", err)
		}
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, len(body["articles"]), len(testArticleData))
	})
	t.Run("Should Return 500 if db returns error", func(t *testing.T) {
		mockDB := ksql.Mock{
			QueryFn: func(ctx context.Context, record interface{}, query string, params ...interface{}) error {
				return errors.New("Something went grown")
			},
		}
		r := setupTestingRouter(mockDB)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/all", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}

// Test the 'Paginated' controller
func Test_PaginatedController(t *testing.T){
  t.Run("Should return 400 if pass invaild page param", func(t *testing.T) {
    r := setupTestingRouter(ksql.Mock{})
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/paginated?page=asdf", nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, 400, w.Code)
  })
}
