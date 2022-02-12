CREATE TABLE users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password TEXT,
    secret_key VARCHAR(255) NOT NULL,
    link_resource TEXT,
    PRIMARY KEY(id),
    UNIQUE (email),
    UNIQUE (secret_key)
);