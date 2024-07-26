package main

import (
	"fmt"

	"github.com/r0m43g/rest-api-demo/internal/comment"
	"github.com/r0m43g/rest-api-demo/internal/db"
	transport "github.com/r0m43g/rest-api-demo/internal/transport/http"
)

// Run is the entry point of our application
func Run() error {
	fmt.Println("Starting up our application...")
	db, err := db.NewDatabase()
	if err != nil {
    return fmt.Errorf("failed to connect to database: %v", err)
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database: ", err)

		return err
	}

	cmtService := comment.NewService(db)

	httpHandler := transport.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

// main is the entry point of our application
func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Go REST API")
}
