-- +goose Up
CREATE TABLE feeds(
    name VARCHAR NOT NULL,
    url  VARCHAR UNIQUE NOT NULL,
    user_id uuid NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
ALTER TABLE feeds DROP CONSTRAINT IF EXISTS fk_user_id;
DROP TABLE feeds;
