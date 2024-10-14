CREATE TABLE tasks(
                       id SERIAL PRIMARY KEY,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       expired_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       description TEXT,
                       user_id SERIAL NOT NULL,
                       FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                       name VARCHAR(50)
)