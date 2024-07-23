package db

import (
  "fmt"
  "github.com/golang-migrate/migrate/v4"
  "github.com/golang-migrate/migrate/v4/database/mysql"
  _ "github.com/golang-migrate/migrate/v4/source/file"
  _ "github.com/go-sql-driver/mysql"
)

func (d *Database) MigrateDB() error {
  fmt.Println("Migrating database...")

  driver, err := mysql.WithInstance(d.Client.DB, &mysql.Config{})
  if err != nil {
    return fmt.Errorf("Failed to create driver: %v", err)
  }

  m, err := migrate.NewWithDatabaseInstance(
    "file://migrations",
    "mysql",
    driver,
  )
  if err != nil {
    fmt.Println(err)
    return err
  }

  if err := m.Up(); err != nil {
    return fmt.Errorf("could not run up migrations: %w" , err)
  }

  fmt.Println("Database migrated successfully!")

  return nil
}
