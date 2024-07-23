package db

import (
  "fmt"
  "context"

  "github.com/jmoiron/sqlx"
  _ "github.com/go-sql-driver/mysql"
)

type Database struct {
  Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
  connectionString := "user:password@/database"

  dbConn, err := sqlx.Connect("mysql", connectionString)
  if err != nil {
    return &Database{}, fmt.Errorf("Failed to connect to database: %v", err)
  }

  return &Database{
    Client: dbConn,
  }, nil
}

func (d *Database) Ping(ctx context.Context) error {
  if err := d.Client.PingContext(ctx); err != nil {
    return fmt.Errorf("Failed to ping database: %v", err)
  }
  return nil
}
