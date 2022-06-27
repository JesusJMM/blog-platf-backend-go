package posts

import "time"

type PartialArticle struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Desc      *string   `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Slug      string    `json:"slug"`
	SmImg     *string   `json:"smImg"`
	LgImg     *string   `json:"lgImg"`
	UserID    int       `json:"userID"`
}

type Author struct {
	ID       int     `json:"userId"`
	Name     string  `json:"name"`
	Img      *string `json:"img"`
}
