CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(2048) NOT NULL
);

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    group_id INT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    title VARCHAR(2048) NOT NULL,
    text TEXT,
    release_date VARCHAR(255),
    link VARCHAR(4000)
);


