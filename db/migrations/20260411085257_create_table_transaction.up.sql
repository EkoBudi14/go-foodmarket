CREATE TABLE transaction
(
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    food_id VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    total INT NOT NULL,
    status VARCHAR(100) NOT NULL,
    payment_url TEXT NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (food_id) REFERENCES food(id)
) ENGINE = InnoDB;