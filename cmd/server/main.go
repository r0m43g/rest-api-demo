package main

import (
	"context"
	"fmt"
	"github.com/r0m43g/rest-api-demo/internal/comment"
	"github.com/r0m43g/rest-api-demo/internal/db"
)

// Run is the entry point of our application
func Run() error {
	fmt.Println("Starting up our application...")
	db, err := db.NewDatabase()
	if err != nil {

		return fmt.Errorf("Failed to connect to database: %v", err)
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database: ", err)

		return err
	}

	cmtService := comment.NewCommentService(db)
	c, err := cmtService.PostComment(context.Background(), comment.Comment{
		Slug:   "my-sudden-post",
		Body:   "This is sudden test post",
		Author: "Romain",
	})
	if err != nil {
		fmt.Println("Error posting comment: ", err)

		return err
	}
	fmt.Println("Comment: ", c)

	c, err = cmtService.GetComment(context.Background(), "ba347452-49c0-11ef-91f2-005056483dcf")
	if err != nil {
		fmt.Println("Error getting comment: ", err)

		return err
	}
	fmt.Println("Comment: ", c)

	return nil
}

// main is the entry point of our application
func main() {

	if err := Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Go REST API")
}
