package auth

import (
	"net/http"

	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/JesusJMM/blog-plat-go/postgres/repos/users"

	"github.com/gin-gonic/gin"
	"github.com/vingarcia/ksql"
	"golang.org/x/net/context"
)

type AuthHandler struct {
	db  ksql.DB
	ctx context.Context
  userRepo users.UserRepository
}

func New(db ksql.DB, ctx context.Context, userRepo users.UserRepository) *AuthHandler {
	return &AuthHandler{
		db,
		ctx,
    userRepo,
	}
}

type SignupPayload struct {
  Name     string `json:"name" binding:"required"`
  Password string `json:"password" binding:"required"`
  Img      string `json:"img"`
}

// Create a new user
// METHOD: POST
func (h AuthHandler) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
    var payload SignupPayload
    if err := c.ShouldBindJSON(&payload); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return 
    }
    newUser, err := h.userRepo.Create(postgres.User{
      Name: payload.Name,
      Password: payload.Password,
      Img: &payload.Img,
    })
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    token, err := SignToken(newUser.ID)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.Header("Authorization", token)
    c.JSON(http.StatusCreated, gin.H{"user": newUser})
  }
}

type LoginPayload struct {
  Name string `json:"name"`
  Password string `json:"password"`
}
// Authenticate user and return token
func (h AuthHandler) Login() gin.HandlerFunc {
  return func(c *gin.Context) {
    var payload LoginPayload
    if err:= c.ShouldBindJSON(&payload); err != nil{
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return 
    }
    var dbUser postgres.User
    err := h.db.QueryOne(h.ctx, &dbUser, "FROM users WHERE name $1 LIMIT 1", payload.Name)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return 
    }
    if !users.ValidPassword(payload.Password, dbUser.Password){ 
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid password"})
      return 
    }
    token, err := SignToken(dbUser.ID)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.Header("Authorization", token)
    c.JSON(http.StatusCreated, gin.H{"user": dbUser})
  }
}
