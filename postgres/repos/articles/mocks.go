package articles

import (
	"time"

	"github.com/JesusJMM/blog-plat-go/postgres"
)

const FakeArticleID = 3

type MockedArticleRepo struct {}

func NewMockedArticleRepo() (r MockedArticleRepo) {
  return
}

func (r MockedArticleRepo) Create(a postgres.Article) (postgres.Article, error){
  a.ID = FakeArticleID
  a.CreatedAt = time.Now()
  a.UpdatedAt = time.Now()
  return a, nil
}

func (r MockedArticleRepo) Update(*UpdateArticleParams) error{
  return nil
}

func (r MockedArticleRepo) Delete(articleID int, authorID int) error {
  return nil
}
