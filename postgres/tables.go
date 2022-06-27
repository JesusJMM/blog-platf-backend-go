package postgres

import(
	"github.com/vingarcia/ksql"
)

var UserTable = ksql.NewTable("users", "user_id")
var ArticleTable = ksql.NewTable("articles", "article_id")
