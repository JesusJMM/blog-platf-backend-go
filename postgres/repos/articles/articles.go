package articles

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/vingarcia/ksql"
)

type ArticleRepository interface {
  Create(postgres.Article) (postgres.Article, error)
  Update(postgres.Article) (postgres.Article, error) 
  Delete(int) (error) 
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
  err = r.db.Transaction(r.ctx, func(p ksql.Provider) error {
    // Check if the user has another article with the same slug
    var conflictArticle postgres.Article
    err := p.QueryOne(
      r.ctx,
      conflictArticle,
      "FROM articles WHERE slug=$1 AND user_id=$2 LIMIT 1",
      outArticle.Slug,
      outArticle.UserID,
    )
    if err != nil {
      if !errors.Is(err, sql.ErrNoRows){
        return fmt.Errorf("Error inserting article: there exist another article with the same slug and user_id")
      }
      return err
    }
    err = p.Insert(r.ctx, postgres.ArticleTable, &outArticle)
    if err != nil{
      return err
    }
    return nil
  })
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

func (r ArticleRepo) Update(article *UpdateArticleParams) (error){
  err := r.db.Transaction(r.ctx, func(p ksql.Provider) error {
    // Check if the user has another article with the same slug
    var targetArticle postgres.Article
    err := p.QueryOne(
      r.ctx,
      &targetArticle,
      "FROM articles WHERE id=$1",
      article.ID,
    )
    var conflictArticle postgres.Article
    err = p.QueryOne(
      r.ctx,
      conflictArticle,
      "FROM articles WHERE slug=$1 AND user_id=$2 AND article_id!=$3 LIMIT 1",
      targetArticle.Slug,
      targetArticle.UserID,
      targetArticle.ID,
    )
    if err != nil {
      if !errors.Is(err, sql.ErrNoRows){
        return fmt.Errorf("Error inserting article: there exist another article with the same slug and user_id")
      }
      return err
    }
    err = p.Patch(r.ctx, postgres.ArticleTable, article)
    if err != nil{
      return err
    }
    return nil
  })
  return err
}

func (r ArticleRepo) Delete(id int) (error) {
  return r.db.Delete(r.ctx, postgres.ArticleTable, id)
}
