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
	Update(*UpdateArticleParams) error
	Delete(int, int) error
}

type ArticleRepo struct {
	db  *ksql.DB
	ctx context.Context
}

func NewArticleRepo(db *ksql.DB, ctx context.Context) ArticleRepo {
	return ArticleRepo{
		db:  db,
		ctx: ctx,
	}
}

var ErrArticleConflict = errors.New("Conflict with another article's slug attribute")
var ErrInvalidAuthor = errors.New("User is not the owner of the resource")

func (r ArticleRepo) Create(article postgres.Article) (outArticle postgres.Article, err error) {
	outArticle = article
  outArticle.CreatedAt = time.Now()
  outArticle.UpdatedAt = time.Now()
	err = r.db.Transaction(r.ctx, func(p ksql.Provider) error {
		// Check if the user has another article with the same slug
		var conflictArticle postgres.Article
		err := p.QueryOne(
			r.ctx,
			&conflictArticle,
			"FROM articles WHERE slug=$1 AND user_id=$2 LIMIT 1",
			outArticle.Slug,
			outArticle.UserID,
		)
		fmt.Println("error: ", err, "conflictArticle: ", conflictArticle)
		if err == nil {
			return ErrArticleConflict
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		err = p.Insert(r.ctx, postgres.ArticleTable, &outArticle)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

type UpdateArticleParams struct {
	ID        int        `ksql:"article_id" json:"id" binding:"required"`
	Title     *string    `ksql:"title" json:"title"`
	Desc      *string    `ksql:"description" json:"desc"`
	Content   *string    `ksql:"content" json:"content"`
	UpdatedAt *time.Time `ksql:"updated_at" json:"-"`
	Slug      *string    `ksql:"slug" json:"slug"`
	SmImg     *string    `ksql:"sm_img" json:"smImg"`
	LgImg     *string    `ksql:"lg_img" json:"lgImg"`
	UserID    int        `ksql:"user_id" json:"-"`
}

func nullable[t any](val t) *t{
  return &val
}

func (r ArticleRepo) Update(article *UpdateArticleParams) error {
  article.UpdatedAt = nullable(time.Now())
	err := r.db.Transaction(r.ctx, func(p ksql.Provider) error {
		// Check if the user has another article with the same slug
		var targetArticle postgres.Article
		err := p.QueryOne(
			r.ctx,
			&targetArticle,
			"FROM articles WHERE article_id=$1 AND user_id=$2",
			article.ID,
			article.UserID,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ErrInvalidAuthor
			}
			return err
		}
		var conflictArticle postgres.Article
		err = p.QueryOne(
			r.ctx,
			&conflictArticle,
			"FROM articles WHERE slug=$1 AND user_id=$2 AND article_id!=$3 LIMIT 1",
			targetArticle.Slug,
			targetArticle.UserID,
			targetArticle.ID,
		)
		fmt.Printf("conflictArticle: %v\n", conflictArticle)
		if err == nil {
			return ErrArticleConflict
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		err = p.Patch(r.ctx, postgres.ArticleTable, article)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r ArticleRepo) Delete(id, userID int) error {
	_, err := r.db.Exec(r.ctx, "DELETE FROM articles WHERE article_id=$1 AND user_id=$2", id, userID)
	return err
}
