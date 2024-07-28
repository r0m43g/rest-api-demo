package comment

import (
	"context"
	"errors"
	"fmt"
)

// Errors
var (
	ErrFetchingComment = errors.New("Failed to fetch comment")
	NotImplemented     = errors.New("Not implemented")
)

// Comment is a struct that represents a comment in our application
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// CommentServiceInterface is an interface that represents a service that manages comments
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
}

// CommentService is a struct that represents a service that manages comments
type Service struct {
	Store Store
}

// NewCommentService creates a new CommentService
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment gets a comment by its ID
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Getting comment with ID: ", id)
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println("Error getting comment: ", err)

		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

// PostComment creates a new comment
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	fmt.Println("Posting comment: ", cmt)
	cmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		fmt.Println("Error posting comment: ", err)

		return Comment{}, err
	}

	return cmt, nil
}

// UpdateComment updates a comment
func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	fmt.Println("Updating comment: ", cmt)
	cmt, err := s.Store.UpdateComment(ctx, id, cmt)
	if err != nil {
		fmt.Println("Error updating comment: ", err)

		return Comment{}, err
	}

	return cmt, nil
}

// DeleteComment deletes a comment
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	fmt.Println("Deleting comment with ID: ", id)
	err := s.Store.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println("Error deleting comment: ", err)

		return err
	}

	return nil
}
