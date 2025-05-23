-- +goose Up
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    hash_id VARCHAR(36) NOT NULL,
    email VARCHAR(255) NOT NULL,
    fullname VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (hash_id),
    UNIQUE (email)
);

-- +goose Down
DROP TABLE users;
