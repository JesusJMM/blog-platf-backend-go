package postgres

import "time"

// "github.com/vingarcia/ksql"

type Post struct {
	ID        int       `ksql:"post_id"`
	Title     string    `ksql:"title"`
	Desc      string    `ksql:"desc"`
	Content   string    `ksql:"content"`
	CreatedAt time.Time `ksql:"created_at"`
	UpdatedAt time.Time `ksql:"updated_at"`
	Slug      string    `ksql:"slug"`
	SmImg     string    `ksql:"sm_img"`
	LgImg     string    `ksql:"lg_img"`
	UserId    string    `ksql:"user_id"`
}

type User struct {
	ID       int    `ksql:"user_id"`
	Name     string `ksql:"name"`
	Password string `ksql:"password"`
	Img      string `ksql:"img"`
}
