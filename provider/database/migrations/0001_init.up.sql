CREATE TABLE IF NOT EXISTS todos(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    todo_value VARCHAR(255),
    is_completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT now(),
    archieved_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS users(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    archieved_at TIMESTAMP DEFAULT NULL
)