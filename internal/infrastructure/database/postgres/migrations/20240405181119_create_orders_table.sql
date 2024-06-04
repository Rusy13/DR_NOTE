-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
                        order_uid VARCHAR PRIMARY KEY,
                        track_number VARCHAR,
                        entry VARCHAR,
                        locale VARCHAR,
                        internal_signature VARCHAR,
                        customer_id VARCHAR,
                        delivery_service VARCHAR,
                        shardkey VARCHAR,
                        sm_id INTEGER,
                        date_created TIMESTAMPTZ,
                        oof_shard VARCHAR
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
