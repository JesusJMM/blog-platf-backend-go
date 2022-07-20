package testu

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// Sends request to router
func MakeRequest(router *gin.Engine, method, url string, body io.Reader) (*httptest.ResponseRecorder){
  w := httptest.NewRecorder()
  req, _ := http.NewRequest(method, url, body)
  router.ServeHTTP(w, req)
  return w
}
