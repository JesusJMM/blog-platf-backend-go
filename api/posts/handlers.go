package posts

import (
	"context"

	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/gin-gonic/gin"
	"github.com/vingarcia/ksql"
)

type PostsHandler struct {
	db  *ksql.DB
	ctx context.Context
}

func New(db *ksql.DB, ctx context.Context) PostsHandler {
	return PostsHandler{
		db:  db,
		ctx: ctx,
	}
}

type PostAndUserRow struct {
	Article postgres.Article `tablename:"a"`
	Author postgres.User     `tablename:"u"`
}

func (h PostsHandler) AllPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var posts []PostAndUserRow
		err := h.db.Query(h.ctx, &posts,
			"FROM articles as a LEFT JOIN users as u on a.user_id = a.user_id",
		)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"posts": posts})
	}
}
