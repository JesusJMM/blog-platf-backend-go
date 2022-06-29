package articles

import (
	"context"
	"time"

	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/vingarcia/ksql"
)

type ArticleRepository interface {
  Create(postgres.Article) (postgres.Article, error)
  Update(postgres.Article) (postgres.Article, error) 
  Delete(postgres.Article) (postgres.Article, error) 
}

type ArticleRepo struct {
  db ksql.DB
  ctx context.Context
}

func NewArticleRepo(db ksql.DB) ArticleRepo {
  return ArticleRepo{
    db: db,
  }
}

func (r ArticleRepo) Create(article postgres.Article) (outArticle postgres.Article, err error){
  outArticle = article
  err = r.db.Insert(r.ctx, postgres.ArticleTable, &outArticle)
  return
}

type UpdateArticleParams struct {
	ID        int       `ksql:"article_id" json:"id"`
	Title     *string   `ksql:"title" json:"title"`
	Desc      *string   `ksql:"description" json:"desc"`
	Content   *string   `ksql:"content" json:"content"`
	UpdatedAt *time.Time `ksql:"updated_at" json:"updated_at"`
	Slug      *string   `ksql:"slug" json:"slug"`
	SmImg     *string   `ksql:"sm_img" json:"smImg"`
	LgImg     *string   `ksql:"lg_img" json:"lgImg"`
	UserID    int       `ksql:"user_id" json:"userID"`
}

func (r ArticleRepo) Update(article UpdateArticleParams) (error){
  return r.db.Patch(r.ctx, postgres.ArticleTable, article)
}
func (r ArticleRepo) Delete(id int) (error) {
  return r.db.Delete(r.ctx, postgres.ArticleTable, id)
}
