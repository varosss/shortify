CREATE TABLE short_links (
    id SERIAL PRIMARY KEY,
    code TEXT UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    user_id INT REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);
