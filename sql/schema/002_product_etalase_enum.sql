-- Write your migrate up statements here

CREATE TYPE product_etalase AS ENUM (
    'recommendation',
    'trending',
    'new',
    'discount'
);

---- create above / drop below ----

DROP TYPE product_etalase;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
