
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE categories ADD COLUMN slug TEXT;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE categories DROP COLUMN slug;
