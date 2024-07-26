package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/r0m43g/rest-api-demo/internal/comment"
)

// CommentService is an interface that represents a service that manages comments
type CommentService interface{
  PostComment(context.Context, comment.Comment) (comment.Comment, error)
  GetComment(context.Context, string) (comment.Comment, error)
  UpdateComment(context.Context, string, comment.Comment) (comment.Comment, error)
  DeleteComment(context.Context, string) error
}

// PostComment creates a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
  var cmt comment.Comment
  if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  cmt, err := h.Service.PostComment(r.Context(), cmt)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  if err := json.NewEncoder(w).Encode(cmt); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
}
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {}
