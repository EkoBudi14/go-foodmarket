CREATE TABLE food 
(
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    ingredients TEXT NOT NULL,
    price INT NOT NULL,
    rate DOUBLE NOT NULL,
    types VARCHAR(255) NOT NULL,
    picture_path TEXT NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
) ENGINE = InnoDB;