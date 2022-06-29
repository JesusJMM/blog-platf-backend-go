package postgres

import (
	"time"
)

// Struct of article
type Article struct {
	ID        int       `ksql:"article_id" json:"id"`
	Title     string    `ksql:"title" json:"title"`
	Desc      *string   `ksql:"description" json:"desc"`
	Content   *string   `ksql:"content" json:"content"`
	CreatedAt time.Time `ksql:"created_at" json:"created_at"`
	UpdatedAt time.Time `ksql:"updated_at" json:"updated_at"`
	Slug      string    `ksql:"slug" json:"slug"`
	SmImg     *string   `ksql:"sm_img" json:"smImg"`
	LgImg     *string   `ksql:"lg_img" json:"lgImg"`
	UserID    int       `ksql:"user_id" json:"userID"`
}

// Struct of user
// ( sensitive information )
type User struct {
	ID       int     `ksql:"user_id"`
	Name     string  `ksql:"name"`
	Password string  `ksql:"password"`
	Img      *string `ksql:"img"`
}
