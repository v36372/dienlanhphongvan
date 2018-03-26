
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE products ADD COLUMN thumbnail TEXT;
ALTER TABLE products ADD COLUMN image01 TEXT;
ALTER TABLE products ADD COLUMN image02 TEXT;
ALTER TABLE products ADD COLUMN image03 TEXT;
ALTER TABLE products ADD COLUMN image04 TEXT;
ALTER TABLE products ADD COLUMN image05 TEXT;
ALTER TABLE products ADD COLUMN image06 TEXT;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE products DROP COLUMN thumbnail;
ALTER TABLE products DROP COLUMN image01;
ALTER TABLE products DROP COLUMN image02;
ALTER TABLE products DROP COLUMN image03;
ALTER TABLE products DROP COLUMN image04;
ALTER TABLE products DROP COLUMN image05;
ALTER TABLE products DROP COLUMN image06;
