CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    currency CHAR(3) NOT NULL,
    spent_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    note TEXT,

    -- Foreign Key from expenses.user_id -> users.id [cite: 59, 62]
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    -- Foreign Key from expenses.category_id -> categories.id [cite: 60, 63]
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT,
    -- Check amount > 0 [cite: 61, 64]
    CHECK (amount > 0)
);

-- Index on user_id [cite: 66]
CREATE INDEX idx_expenses_user_id ON expenses (user_id);

-- Composite index on (user_id, spent_at) for range queries [cite: 67]
CREATE INDEX idx_expenses_user_spent ON expenses (user_id, spent_at);
