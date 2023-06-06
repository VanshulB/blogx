// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: post.sql

package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createPost = `-- name: CreatePost :one

INSERT into
    posts (id, user_id, title, body, tags)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, title, body, likes, views, tags, created_at, updated_at
`

type CreatePostParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Title  string
	Body   string
	Tags   []string
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.UserID,
		arg.Title,
		arg.Body,
		pq.Array(arg.Tags),
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.Likes,
		&i.Views,
		pq.Array(&i.Tags),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec

DELETE FROM posts WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const disLikePost = `-- name: DisLikePost :one

DELETE FROM user_likes
WHERE post_id = $1 AND user_id = $2
RETURNING user_id, post_id
`

type DisLikePostParams struct {
	PostID uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DisLikePost(ctx context.Context, arg DisLikePostParams) (UserLike, error) {
	row := q.db.QueryRowContext(ctx, disLikePost, arg.PostID, arg.UserID)
	var i UserLike
	err := row.Scan(&i.UserID, &i.PostID)
	return i, err
}

const likePost = `-- name: LikePost :one

INSERT INTO user_likes(user_id, post_id) VALUES($1, $2) RETURNING user_id, post_id
`

type LikePostParams struct {
	UserID uuid.UUID
	PostID uuid.UUID
}

func (q *Queries) LikePost(ctx context.Context, arg LikePostParams) (UserLike, error) {
	row := q.db.QueryRowContext(ctx, likePost, arg.UserID, arg.PostID)
	var i UserLike
	err := row.Scan(&i.UserID, &i.PostID)
	return i, err
}

const updatePost = `-- name: UpdatePost :one

UPDATE posts
SET
    title = $2,
    body = $3,
    tags = $4
WHERE id = $1
RETURNING id, user_id, title, body, likes, views, tags, created_at, updated_at
`

type UpdatePostParams struct {
	ID    uuid.UUID
	Title string
	Body  string
	Tags  []string
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.ID,
		arg.Title,
		arg.Body,
		pq.Array(arg.Tags),
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.Likes,
		&i.Views,
		pq.Array(&i.Tags),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
