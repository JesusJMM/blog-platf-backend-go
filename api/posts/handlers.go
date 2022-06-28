package posts

import (
	"context"
	"strconv"

	// "github.com/JesusJMM/blog-plat-go/postgres"
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

type PartialPostWithAuthor struct {
	Article PartialArticle `tablename:"a"`
	Author  Author         `tablename:"u"`
}

const PartialArticleQuery = `
    FROM articles as a
    LEFT JOIN users as u
    ON a.user_id = u.user_id
`

const PaginationSize = 10

// Returns all posts in the database
func (h PostsHandler) All() gin.HandlerFunc {
	return func(c *gin.Context) {
		var posts []PartialPostWithAuthor
		err := h.db.Query(h.ctx, &posts,
			PartialArticleQuery,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"posts": posts})
	}
}

// Return a set of PartialPostWithAuthor 
// (paginated route)
func (h PostsHandler) Paginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryPage := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(queryPage)
		if err != nil {
			c.JSON(500, gin.H{"error": "'page' query param must be a number"})
			return
		}
		var posts []PartialPostWithAuthor
    q := PartialArticleQuery + `ORDER BY a.article_id LIMIT $1 OFFSET $2`
		err = h.db.Query(
      h.ctx, 
      &posts,
      q,
      PaginationSize,
      (page -1) * PaginationSize,
		)
    if err != nil {
      c.JSON(500, gin.H{"error": err.Error()})
      return
    }
    c.JSON(200, gin.H{"posts": posts})
	}
}

// Return a set of PartialPostWithAuthor by author
// requires an 'author' url param
// (paginated route)
func (h PostsHandler) ByAuthorPaginated() gin.HandlerFunc {
  return func(c *gin.Context) {
    author := c.Param("author")
		queryPage := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(queryPage)
		if err != nil {
			c.JSON(500, gin.H{"error": "'page' query param must be a number"})
			return
		}
    posts := []PartialPostWithAuthor{}
    err = h.db.Query(h.ctx, &posts,
      PartialArticleQuery + `WHERE u.name=$1 ORDER BY a.article_id LIMIT $2 OFFSET $3`,
      author,
      PaginationSize,
      (page -1) * PaginationSize,
    )
    if err != nil {
      c.JSON(500, gin.H{"error": err.Error()})
      return
    }
    c.JSON(200, gin.H{"posts": posts})
  }
}
