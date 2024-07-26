package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/r0m43g/rest-api-demo/internal/comment"
  validator "github.com/go-playground/validator/v10"
)

// CommentService is an interface that represents a service that manages comments
type CommentService interface{
  PostComment(context.Context, comment.Comment) (comment.Comment, error)
  GetComment(context.Context, string) (comment.Comment, error)
  UpdateComment(context.Context, string, comment.Comment) (comment.Comment, error)
  DeleteComment(context.Context, string) error
}

type PostCommentRequest struct {
  Slug string `json:"slug" validate:"required"`
  Body string `json:"body" validate:"required"`
  Author string `json:"author" validate:"required"`
}

func convertToComment(cmt PostCommentRequest) comment.Comment {
  return comment.Comment{
    Slug: cmt.Slug,
    Body: cmt.Body,
    Author: cmt.Author,
  }
}

// PostComment creates a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
  var cmt PostCommentRequest
  if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  validate := validator.New()
  if err := validate.Struct(cmt); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  convertedCmt := convertToComment(cmt)
  posted, err := h.Service.PostComment(r.Context(), convertedCmt)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := json.NewEncoder(w).Encode(posted); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
}

// GetComment gets a comment by its ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
  id := mux.Vars(r)["id"]
  if id == "" {
    http.Error(w, "missing id in request", http.StatusBadRequest)
    return
  }
  cmt, err := h.Service.GetComment(r.Context(), id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := json.NewEncoder(w).Encode(cmt); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

// UpdateComment updates a comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
  id := mux.Vars(r)["id"]
  if id == "" {
    http.Error(w, "missing id in request", http.StatusBadRequest)
    return
  }
  var cmt comment.Comment
  if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := json.NewEncoder(w).Encode(cmt); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

// DeleteComment deletes a comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
  id := mux.Vars(r)["id"]
  if id == "" {
    http.Error(w, "missing id in request", http.StatusBadRequest)
    return
  }
  if err := h.Service.DeleteComment(r.Context(), id); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusNoContent)
}
