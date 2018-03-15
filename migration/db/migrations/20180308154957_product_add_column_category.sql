
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE products ADD COLUMN category_id INTEGER;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE products DROP COLUMN category_id;

