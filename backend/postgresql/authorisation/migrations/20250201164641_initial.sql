-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION pgcrypto;

CREATE TABLE customers (
    customer_id VARCHAR(255) DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY(customer_id)
);

CREATE TYPE stream_type AS ENUM (
    'SPOTIFY',
    'TIDAL'
);

CREATE TABLE stream_creds (
    customer_id VARCHAR(255),
    PRIMARY KEY(customer_id),
    stream_type stream_type NOT NULL,
    CONSTRAINT fk_customer
        FOREIGN KEY(customer_id)
            REFERENCES customers(customer_id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EXTENSION pgcrypto;
DROP TABLE customers;
DROP TABLE stream_creds;
-- +goose StatementEnd
