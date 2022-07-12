package auth

import (
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
    assert.NotNil(t, err, "Fail to sing token: %s", err)
  })
}

func Test_ParseToken(t *testing.T){
  t.Run("Should return error with invalid token", func(t *testing.T){
    _, _, err := ParseToken("invalid token")
    assert.NotNil(t, err, "Not return error with invalid token")
  })
}

func Test_SingAndParseToken(t *testing.T){
  t.Run("Signed tokens should be parsed", func(t *testing.T){
    tokenString, err := SignToken(testUser)

    assert.NotNilf(t, err, "Fail to sing token: %s", err)

    _, claims, err := ParseToken(tokenString)
    assert.NotNilf(t, err, "Fail to parse token: %s", err)

    assert.Equal(t, claims.UID, testUser.ID)
    assert.Equal(t, claims.UserName, testUser.Name)
  })  

}
