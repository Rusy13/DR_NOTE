-- +goose Up
-- +goose StatementBegin

CREATE TABLE subscriptions (
                               id SERIAL PRIMARY KEY,
                               user_id INT REFERENCES users(id) ON DELETE CASCADE,
                               subscribed_to_id INT REFERENCES users(id) ON DELETE CASCADE,
                               UNIQUE (user_id, subscribed_to_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE subscriptions;
-- +goose StatementEnd
