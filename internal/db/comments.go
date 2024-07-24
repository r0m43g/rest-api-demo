package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/r0m43g/rest-api-demo/internal/comment"
	"github.com/satori/go.uuid"
)

// CommentRow represents a row in the comments table
type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

// GetComment gets a comment by its ID
func convertCommentRowToComment(c CommentRow) comment.Comment {

	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Author.String,
	}
}

// GetComment gets a comment by its ID
func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(ctx, "SELECT id, slug, body, author FROM comments WHERE id = ?", uuid)
	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {

		return comment.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

// PostComment creates a new comment
func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(ctx, "INSERT INTO comments (id, slug, body, author) VALUES (:id, :slug, :body, :author)", postRow)
	if err != nil {

		return comment.Comment{}, fmt.Errorf("error inserting comment: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&postRow.ID, &postRow.Slug, &postRow.Body, &postRow.Author)
		if err != nil {

			return comment.Comment{}, fmt.Errorf("error scanning comment: %w", err)
		}
	}

	return convertCommentRowToComment(postRow), nil
}

// UpdateComment updates a comment
func (d *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(ctx, "UPDATE comments SET slug = :slug, body = :body, author = :author WHERE id = :id", cmtRow)
	if err != nil {

		return comment.Comment{}, fmt.Errorf("error updating comment: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
		if err != nil {

			return comment.Comment{}, fmt.Errorf("error scanning comment: %w", err)
		}
	}

	return convertCommentRowToComment(cmtRow), nil
}

// DeleteComment deletes a comment
func (d *Database) DeleteComment(ctx context.Context, uuid string) error {
	_, err := d.Client.ExecContext(ctx, "DELETE FROM comments WHERE id = ?", uuid)
	if err != nil {

		return fmt.Errorf("error deleting comment: %w", err)
	}

	return nil
}
