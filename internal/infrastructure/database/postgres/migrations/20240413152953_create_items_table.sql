-- +goose Up
-- +goose StatementBegin
CREATE TABLE items (
                       order_uid VARCHAR REFERENCES orders(order_uid),
                       chrt_id INTEGER,
                       track_number VARCHAR,
                       price INTEGER,
                       rid VARCHAR,
                       name VARCHAR,
                       sale INTEGER,
                       size VARCHAR,
                       total_price INTEGER,
                       nm_id INTEGER,
                       brand VARCHAR,
                       status INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd
