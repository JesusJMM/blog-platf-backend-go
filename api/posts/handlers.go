package posts

import (
	"context"

	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/gin-gonic/gin"
	"github.com/vingarcia/ksql"
)

type PostsHandler struct{
  db ksql.DB
  ctx context.Context
}

type PostAndUserRow struct {
  Post postgres.Article `tablename:"p"`
  User postgres.User `tablename:"u"`
}

func (h PostsHandler) AllPosts() gin.HandlerFunc{
  return func(c *gin.Context) {
    var posts PostAndUserRow
    q := `FROM posts as p LEFT JOIN user as u ON posts.user_id = u.user_id LIMIT 300`
    err := h.db.Query(h.ctx, posts, q)
    if err != nil {
      c.JSON(500, gin.H{"error": err.Error()})
    }
    c.JSON(200, gin.H{"posts": posts})
  }
}
