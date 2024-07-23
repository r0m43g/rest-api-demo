package main

import (
  "fmt"
  "github.com/r0m43g/rest-api-demo/internal/db"
)

// Run is the entry point of our application
func Run() error {
  fmt.Println("Starting up our application...")
  db, err := db.NewDatabase()

  if err != nil {
    return fmt.Errorf("Failed to create database: %v", err)
  }

  if err := db.MigrateDB(); err != nil {
    fmt.Println("Failed to migrate database: ", err)
    return err
  }
  fmt.Println("Database migrated successfully!")

  return nil
}

// main is the entry point of our application
func main() {

  if err:= Run(); err != nil {
    fmt.Println(err)
  }

  fmt.Println("Go REST API")
}
