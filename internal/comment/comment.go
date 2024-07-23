package comment

import (
  "context"
  "errors"
  "fmt"
)

// Errors
var (
  ErrFetchingComment = errors.New("Failed to fetch comment")
  NotImplemented = errors.New("Not implemented")
)

// Comment is a struct that represents a comment in our application
type Comment struct {
  ID string
  Slug string
  Body string
  Author string
}

// CommentServiceInterface is an interface that represents a service that manages comments
type Store interface {
  GetComment(context.Context, string) (Comment, error)
  CreateComment(context.Context, Comment) (Comment, error)
  UpdateComment(context.Context, Comment) (Comment, error)
  DeleteComment(context.Context, string) error
}

// CommentService is a struct that represents a service that manages comments
type CommentService struct {
  Store Store
}

// NewCommentService creates a new CommentService
func NewCommentService(store Store) *CommentService {
  return &CommentService{
    Store: store,
  }
}

// GetComment gets a comment by its ID
func (s *CommentService) GetComment(ctx context.Context, id string) (Comment, error) {
  fmt.Println("Getting comment with ID: ", id)
  cmt, err := s.Store.GetComment(ctx, id)
  if err != nil {
    fmt.Println("Error getting comment: ", err)
    return Comment{}, ErrFetchingComment
  }
  return cmt, nil
}

func (s *CommentService) CreateComment(ctx context.Context, cmt Comment) (Comment, error){
  return Comment{}, NotImplemented
}

func (s *CommentService) UpdateComment(ctx context.Context, cmt Comment) (Comment, error) {
  return Comment{}, NotImplemented
}

func (s *CommentService) DeleteComment(ctx context.Context, id string) error {
  return NotImplemented
}
