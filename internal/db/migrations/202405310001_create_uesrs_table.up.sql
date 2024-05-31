CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name  VARCHAR(255) NOT NULL,
    created_at DATETIME(6)  NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6)  NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    deleted_at DATETIME(6)
);

INSERT INTO users (first_name, last_name) VALUES ('Taro', 'Yamada');
INSERT INTO users (first_name, last_name) VALUES ('Hanako', 'Yamada');
