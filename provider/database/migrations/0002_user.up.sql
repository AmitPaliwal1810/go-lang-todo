CREATE TABLE IF NOT EXISTS users(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255),
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    is_archieved TIMESTAMP DEFAULT NULL
)