package auth

import (
	"fmt"
	"testing"

	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/stretchr/testify/assert"
)

var testUser = postgres.User{
      Name: "TestUser",
      ID: 3,
} 

func Test_Signtoken(t *testing.T){
  t.Run("Should not return error with the right arguments", func(t *testing.T){
    img := "https://yt3.ggpht.com/yti/APfAmoFNM8V6_IAI_SPgAIXsH04Zo3tctEgZnCEnwuqa=s88-c-k-c0x00ffffff-no-rj-mo"
    _, err := SignToken(postgres.User{
      Name: "TestUser",
      ID: 3,
      Img: &img,
    })
    assert.Nil(t, err, "Fail to sing token: %s", err)
  })
}

func Test_ParseToken(t *testing.T){
  t.Run("Should return error with invalid token", func(t *testing.T){
    _, _, err := ParseToken("invalidtoken")
    fmt.Printf("error: %v", err)
    if err == nil {
      t.Fatal("Error is nil")
    }
  })
}

func Test_SingAndParseToken(t *testing.T){
  tokenString, err := SignToken(testUser)

  assert.Nilf(t, err, "Fail to sing token: %s", err)

  _, claims, err := ParseToken(tokenString)
  assert.Nilf(t, err, "Fail to parse token: %s", err)

  assert.Equal(t, claims.UID, testUser.ID)
  assert.Equal(t, claims.UserName, testUser.Name)
}
