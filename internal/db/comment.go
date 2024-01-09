package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/anil1226/go-banking/internal/comment"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtrow CommentRow
	row := d.client.QueryRowContext(ctx, `Select id, slug,body, author from comments where id= $1`, uuid)

	err := row.Scan(&cmtrow.ID, &cmtrow.Slug, &cmtrow.Body, &cmtrow.Author)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching the comment by uid: %w", err)
	}

	return convertCommentRowToComment(cmtrow), nil
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Author: c.Author.String,
		Body:   c.Body.String,
	}
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()

	posrRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	rows, err := d.client.NamedQueryContext(
		ctx,
		`insert into comments(id, slug, author, body)
		values
		(:id, :slug, :author, :body)`,
		posrRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return cmt, nil
}

func (d *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {

	posrRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	rows, err := d.client.NamedQueryContext(
		ctx,
		`update comments
		set
		slug= :slug
		author = :author
		body= :body
		where id = :id`,
		posrRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to update comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertCommentRowToComment(posrRow), nil
}

func (d *Database) DeleteComment(ctx context.Context, id string) error {

	_, err := d.client.ExecContext(ctx, `delete from comments where id= $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting the comment by uid: %w", err)
	}

	return nil
}
