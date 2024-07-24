package db

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Database is a struct that represents a database
type Database struct {
	Client *sqlx.DB
}

// NewDatabase creates a new database
func NewDatabase() (*Database, error) {
	connectionString := "user:password@/comdb"
	dbConn, err := sqlx.Connect("mysql", connectionString)
	if err != nil {

		return &Database{}, fmt.Errorf("Failed to connect to database: %v", err)
	}

	return &Database{
		Client: dbConn,
	}, nil
}

// Ping pings the database
func (d *Database) Ping(ctx context.Context) error {
	if err := d.Client.PingContext(ctx); err != nil {
		return fmt.Errorf("Failed to ping database: %v", err)
	}

	return nil
}
