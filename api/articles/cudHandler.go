package articles

import "github.com/gin-gonic/gin"

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
	}
}
