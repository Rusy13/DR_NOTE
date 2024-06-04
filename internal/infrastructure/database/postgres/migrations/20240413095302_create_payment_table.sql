-- +goose Up
-- +goose StatementBegin

CREATE TABLE payment (
                         order_uid VARCHAR REFERENCES orders(order_uid),
                         transaction VARCHAR,
                         request_id VARCHAR,
                         currency VARCHAR,
                         provider VARCHAR,
                         amount INTEGER,
                         payment_dt BIGINT,
                         bank VARCHAR,
                         delivery_cost INTEGER,
                         goods_total INTEGER,
                         custom_fee INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "payment"
-- +goose StatementEnd
