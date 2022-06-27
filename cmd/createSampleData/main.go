package main

import (
	"context"
	"fmt"

	"github.com/JesusJMM/blog-plat-go/postgres"
	"github.com/vingarcia/ksql"
)

var UserTable = ksql.NewTable("users", "user_id")     
var ArticlesTable = ksql.NewTable("articles", "article_id")

func main() {
  db, err := postgres.New()
  handleError(err)
  ctx := context.Background()

  user := createFailUser()
  err = db.Insert(ctx, UserTable, &user)
  handleError(err)
  fmt.Printf("user: %#v", user)

  posts := createFailPosts(user.ID)
  for i, p := range posts {
    err = db.Insert(ctx, ArticlesTable, p)
    handleError(err)
    fmt.Println("Inserted article #",i)
  }
}

func handleError(err error) {
  if err != nil {
    panic(err)
  }
}

