-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       birthday DATE NOT NULL,
                       password TEXT NOT NULL,
                       api_id INT NOT NULL,
                       api_hash VARCHAR(255) NOT NULL,
                       phone VARCHAR(20) NOT NULL
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
