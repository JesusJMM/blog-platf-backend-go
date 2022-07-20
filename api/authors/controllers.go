package authors

import (
	"context"
	"log"
	"net/http"

	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/gin-gonic/gin"
	"github.com/vingarcia/ksql"
)

type AuthorsService struct {
  db ksql.Provider
  ctx context.Context
}

func New(db ksql.Provider, ctx context.Context) *AuthorsService {
  return &AuthorsService{
    db,
    ctx,
  }
}

func (s AuthorsService) GetAll() gin.HandlerFunc {
  return func(c *gin.Context){
    var users = []postgres.User{}
    q := `FROM users`
    err := s.db.Query(s.ctx, &users, q)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": "Internal Server Error",
      })
    }
    c.JSON(http.StatusOK, gin.H{"authors": users})
  }
}
func (s AuthorsService) One() gin.HandlerFunc {
  return func(c *gin.Context){
    var users = postgres.User{}
    q := `FROM users WHERE users.name=$1 LIMIT 1`
    err := s.db.QueryOne(s.ctx, &users, q, c.Param("authorName"))
    if err != nil {
      log.Println("err: ",err)
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": "Internal Server Error",
      })
      return
    }
    c.JSON(http.StatusOK, gin.H{"author": users})
  }
}
