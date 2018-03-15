
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE categories (
	id SERIAL PRIMARY KEY, 
	name TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE categories;
