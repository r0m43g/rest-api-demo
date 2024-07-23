package main

import (
  "fmt"
  "github.com/r0m43g/rest-api-demo/db"
)

func Run() error {
  fmt.Println("Starting up our application...")
  db, err := db.NewDatabase()
  if err != nil {
    return fmt.Errorf("Failed to create database: %v", err)
  }
  if err := db.Ping(); err != nil {
    return fmt.Errorf("Failed to ping database: %v", err)
  }

  return nil
}

func main() {
  if err:= Run(); err != nil {
    fmt.Println(err)
  }
  fmt.Println("Go REST API")
}
