-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION pgcrypto;

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
    customer_id uuid PRIMARY KEY,
    auth_code VARCHAR(255) NOT NULL,
    stream_type stream_type NOT NULL,
    CONSTRAINT fk_customer
        FOREIGN KEY(customer_id)
            REFERENCES customers(id)
            ON DELETE CASCADE
);

SELECT 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EXTENSION pgcrypto;
DROP TABLE stream_creds;
DROP TYPE stream_type;
DROP TABLE customers;
-- +goose StatementEnd
