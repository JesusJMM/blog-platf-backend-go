package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

// Return an instance of ksql.DB
// uses POSTGRES_URL enviroment variable
func New() (*ksql.DB, error) {
	ctx := context.Background()

	db, err := kpgx.New(ctx, os.Getenv("POSTGRES_URL"), ksql.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %w", err)
	}
	return &db, nil
}
