-- +goose Up
-- +goose StatementBegin
CREATE TABLE customers (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TYPE stream_type AS ENUM (
    'SPOTIFY',
    'TIDAL'
);

CREATE TABLE stream_creds (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id uuid NOT NULL,
    auth_code VARCHAR(255) NOT NULL,
    stream_type stream_type NOT NULL,
    UNIQUE(customer_id, stream_type),
    CONSTRAINT fk_customer
        FOREIGN KEY(customer_id)
            REFERENCES customers(id)
            ON DELETE CASCADE
);

CREATE TABLE token_store (
    token VARCHAR(128) PRIMARY KEY,
    customer_id uuid NOT NULL,
    expire_after DATE NOT NULL,
    CONSTRAINT fk_customer
        FOREIGN KEY(customer_id)
            REFERENCES customers(id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE token_store;
DROP TABLE stream_creds;
DROP TYPE stream_type;
DROP TABLE customers;
-- +goose StatementEnd
