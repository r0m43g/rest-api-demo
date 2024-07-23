package main

import (
  "fmt"
  "context"
  "github.com/r0m43g/rest-api-demo/internal/db"
)

func Run() error {
  fmt.Println("Starting up our application...")
  db, err := db.NewDatabase()
  if err != nil {
    return fmt.Errorf("Failed to create database: %v", err)
  }
  if err := db.Ping(context.Background()); err != nil {
    return fmt.Errorf("Failed to ping database: %v", err)
  }
  fmt.Println("Database is up and running!")
  return nil
}

func main() {
  if err:= Run(); err != nil {
    fmt.Println(err)
  }
  fmt.Println("Go REST API")
}
