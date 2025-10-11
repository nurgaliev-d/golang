CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    user_id INTEGER,
    -- Foreign key constraint: references users(id) if not null [cite: 37]
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    -- Unique constraint: category name must be unique per user [cite: 38, 39]
    UNIQUE (user_id, name)
);

-- Index on user_id [cite: 41]
CREATE INDEX idx_categories_user_id ON categories (user_id);
