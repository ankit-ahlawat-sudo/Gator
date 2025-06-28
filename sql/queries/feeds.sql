-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedsFromDb :many
Select feeds.name AS feed_name, feeds.url, users.name AS user_name FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: GetFeedFromName :one
SELECT * from feeds WHERE feeds.name = $1;

-- name: GetFeedFromURL :one
SELECT * from feeds WHERE feeds.url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds 
SET last_fetched_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT url from feeds 
ORDER BY last_fetched_at ASC NULLS FIRST 
LIMIT 1;
