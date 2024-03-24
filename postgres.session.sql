CREATE TABLE IF NOT EXISTS todos(
            id  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
            todoValue VARCHAR(200),
            iscompleted BOOLEAN DEFAULT FALSE,
            createAt TIMESTAMP DEFAULT now(),
            isArchieved TIMESTAMP DEFAULT  now()
);

INSERT INTO todos (todoValue) VALUES ('Need to complete the task for the new admin UI');