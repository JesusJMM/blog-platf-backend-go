package articles

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	// "github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/JesusJMM/blog-plat-go/postgres/repos/articles"
	"github.com/gin-gonic/gin"
	"github.com/vingarcia/ksql"
)

type ArticleHandler struct {
	db  ksql.Provider 
	ctx context.Context
  articleRepo articles.ArticleRepository
}

func New(db ksql.Provider, ctx context.Context, articleRepo articles.ArticleRepository) ArticleHandler {
	return ArticleHandler{
		db:  db,
		ctx: ctx,
    articleRepo: articleRepo,
	}
}

type PartialArticleWithAuthor struct {
  Article PartialArticle `tablename:"a" json:"article"`
  Author  postgres.User  `tablename:"u" json:"author"`
}

const PartialArticleQuery = `
    FROM articles as a
    LEFT JOIN users as u
    ON a.user_id = u.user_id
`

const PaginationSize = 10

// Returns all articles in the database
// METHOD: GET
func (h ArticleHandler) All() gin.HandlerFunc {
	return func(c *gin.Context) {
		var articles []PartialArticleWithAuthor
		err := h.db.Query(h.ctx, &articles,
			PartialArticleQuery,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"articles": articles})
	}
}

// Return a set of PartialPostWithAuthor
// METHOD: GET (paginated route)
func (h ArticleHandler) Paginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryPage := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(queryPage)
		if err != nil {
			c.JSON(500, gin.H{"error": "'page' query param must be a number"})
			return
		}
		var articles []PartialArticleWithAuthor
		err = h.db.Query(
			h.ctx,
			&articles,
			PartialArticleQuery + `ORDER BY a.article_id DESC LIMIT $1 OFFSET $2`,
			PaginationSize,
			(page-1)*PaginationSize,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"articles": articles})
	}
}

// Return a set of PartialAritcleWithAuthor by author
// requires an 'author' url param
// METHOD: GET (paginated route)
func (h ArticleHandler) ByAuthorPaginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		author := c.Param("author")
		queryPage := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(queryPage)
		if err != nil {
			c.JSON(500, gin.H{"error": "'page' query param must be a number"})
			return
		}
		articles := []PartialArticleWithAuthor{}
		err = h.db.Query(h.ctx, &articles,
			PartialArticleQuery+`WHERE u.name=$1 ORDER BY a.article_id DESC LIMIT $2 OFFSET $3`,
			author,
			PaginationSize,
			(page-1)*PaginationSize,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"articles": articles})
	}
}

func (h ArticleHandler) OneArticle() gin.HandlerFunc {
  return func(c *gin.Context) {
    slug := c.Param("slug")
    author := c.Param("author")
    article := struct {
      Article postgres.Article `tablename:"a" json:"article"`
      Author postgres.User `tablename:"u" json:"author"`
    }{}
    err := h.db.QueryOne(h.ctx, &article, 
      PartialArticleQuery+`WHERE a.slug=$1 AND u.name=$2 LIMIT 1`,
      slug,
      author,
    )
    if err != nil {
      if errors.Is(err, sql.ErrNoRows) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
        return
      }
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, article)
  }
}
