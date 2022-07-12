package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

// Return an instance of ksql.DB
// uses POSTGRES_URL enviroment variable
func New() (*ksql.DB, error) {
	ctx := context.Background()
  if os.Getenv("POSTGRES_URL") == "" {
    log.Fatal("POSTGRES_URL enviroment variable is not set.")    
  }
	db, err := kpgx.New(ctx, os.Getenv("POSTGRES_URL"), ksql.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %w", err)
	}
	return &db, nil
}
