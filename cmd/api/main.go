package main

import (
	"log"

	"github.com/JesusJMM/blog-plat-go/api"
	"github.com/JesusJMM/blog-plat-go/postgres"
)

func main() {
	db, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}
	r := api.New(db)
	r.Run()
}
