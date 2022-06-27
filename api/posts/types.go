package posts

import "time"

type PartialArticle struct {
	ID        int       `ksql:"article_id" json:"id"`
	Title     string    `ksql:"title" json:"title"`
	Desc      *string   `ksql:"description" json:"desc"`
	CreatedAt time.Time `ksql:"created_at" json:"created_at"`
	UpdatedAt time.Time `ksql:"updated_at" json:"updated_at"`
	Slug      string    `ksql:"slug" json:"slug"`
	SmImg     *string   `ksql:"sm_img" json:"smImg"`
	LgImg     *string   `ksql:"lg_img" json:"lgImg"`
	UserID    int       `ksql:"user_id" json:"userID"`
}

type Author struct {
	ID   int     `ksql:"user_id" json:"userId"`
	Name string  `ksql:"name" json:"name"`
	Img  *string `ksql:"img" json:"img"`
}
