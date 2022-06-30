package articles

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/JesusJMM/blog-plat-go/postgres/repos/articles"
	"github.com/gin-gonic/gin"
)

type CreatePayload struct {
	Title   string `json:"title" bindind:"required"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	Slug    string `json:"slug" bindind:"required"`
	SmImg   string `json:"smImg"`
	LgImg   string `json:"lgImg"`
}

// Create a new Article
// Requires authorization
// METHOD: POST
func (h ArticleHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
    var payload CreatePayload
    if err := c.ShouldBindJSON(payload); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    dbArticle, err := h.articleRepo.Create(postgres.Article{
      Title: payload.Title,
      Desc: &payload.Desc,
      Content: &payload.Content,
      Slug: payload.Slug,
      SmImg: &payload.SmImg,
      LgImg: &payload.LgImg,
    })
    if err != nil {
      if errors.Is(err, articles.ErrArticleConflict){
        c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
      }
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusCreated, gin.H{"article": dbArticle})
	}
}

func (h ArticleHandler) Update() gin.HandlerFunc {
  return func(c *gin.Context) {
    var payload articles.UpdateArticleParams
    if err := c.ShouldBindJSON(payload); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
    err := h.articleRepo.Update(&payload)
    if err != nil{
      if errors.Is(err, articles.ErrArticleConflict){
        c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
      }
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.String(http.StatusOK, "Created")
  }
}

func (h ArticleHandler) Delete() gin.HandlerFunc {
  return func(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil{
      c.JSON(http.StatusBadRequest, gin.H{"error": "id url params is not a valid integer"})
      return
    }
    h.articleRepo.Delete(id)
    c.Status(http.StatusOK)
  }
}
