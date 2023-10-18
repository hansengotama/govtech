-- Write your migrate up statements here

CREATE TABLE product_reviews (
    id serial PRIMARY KEY,
    product_id integer NOT NULL REFERENCES products(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    review text NOT NULL,
    rating integer NOT NULL,

    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);

---- create above / drop below ----

drop table product_reviews;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
