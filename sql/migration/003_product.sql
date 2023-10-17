-- Write your migrate up statements here

CREATE TABLE products (
    id serial PRIMARY KEY,
    category_id integer NOT NULL REFERENCES product_categories(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    sku text UNIQUE NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    etalase product_etalase[] DEFAULT ARRAY[]::product_etalase[],
    images jsonb[] DEFAULT ARRAY[]::jsonb[],
    weight real NOT NULL,
    price real NOT NULL,
    rating real,

    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);

---- create above / drop below ----

drop table products;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
