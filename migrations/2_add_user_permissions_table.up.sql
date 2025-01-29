CREATE TABLE user_permissions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    permission TEXT,

    FOREIGN KEY (user_id) REFERENCES users(id)
);