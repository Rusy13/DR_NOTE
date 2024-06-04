-- +goose Up
-- +goose StatementBegin

CREATE TABLE delivery (
                          order_uid VARCHAR REFERENCES orders(order_uid),
                          name VARCHAR,
                          phone VARCHAR,
                          zip VARCHAR,
                          city VARCHAR,
                          address VARCHAR,
                          region VARCHAR,
                          email VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE delivery;
-- +goose StatementEnd
