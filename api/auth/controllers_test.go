package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/JesusJMM/blog-plat-go/lib/testu"
	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/JesusJMM/blog-plat-go/postgres/repos/users"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vingarcia/ksql"
)

func setupTestingRouter(db ksql.Provider, userRepo users.UserRepository) *gin.Engine {
  controllers := New(db, context.Background(), userRepo)
  gin.SetMode(gin.TestMode)
  r := gin.New()
  r.POST("/signup", controllers.Signup())
  r.POST("/login", controllers.Login())
  return r
}

type commonBodyType map[string][]map[string]interface{}

func Test_SignupController(t *testing.T){
  testUser := map[string]interface{}{
    "name": "jhon",
    "password": "password123",
    "img": "https://avatars.githubusercontent.com/u/66509065?v=4",
  }

  t.Run("Should register the user", func(t *testing.T) {
    mockUserRepo := users.MockedUserRepo{
      CreateFn: func(user postgres.User) (postgres.User, error){
        user.ID = 3
        return user, nil
      },
    }
    body, _ := json.Marshal(testUser)
    reader := bytes.NewReader(body)
    w := testu.MakeRequest(
      setupTestingRouter(ksql.Mock{}, mockUserRepo),
      "POST",
      "/signup",
      reader,
      )
    assert.Equal(t, 201, w.Code)
  })
  t.Run("Should return 403 if user already exist", func(t *testing.T) {
    mockUserRepo := users.MockedUserRepo{
      CreateFn: func(user postgres.User) (postgres.User, error){
        user.ID = 3
        return user, errors.New(`ERROR: duplicate key value violates unique constraint \"unique_name\" (SQLSTATE 23505)`)
      },
    }
    body, _ := json.Marshal(testUser)
    reader := bytes.NewReader(body)
    w := testu.MakeRequest(
      setupTestingRouter(ksql.Mock{}, mockUserRepo),
      "POST",
      "/signup",
      reader,
      )
    assert.Equal(t, http.StatusConflict, w.Code)
  })
}
