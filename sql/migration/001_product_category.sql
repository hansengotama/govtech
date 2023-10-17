-- Write your migrate up statements here

CREATE TABLE product_categories (
    id serial PRIMARY KEY,
    guid uuid NOT NULL DEFAULT gen_random_uuid(),
    name text UNIQUE NOT NULL,

    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);

---- create above / drop below ----

drop table product_categories;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
