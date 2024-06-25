-- runners
CREATE TABLE IF NOT EXISTS runners (
    id INTEGER PRIMARY KEY NOT NULL,
    name VARCHAR(60) NOT NULL,
    age INTEGER,
    country VARCHAR(60) NOT NULL,
    season_best TEXT NOT NULL,
    personal_best TEXT NOT NULL
);
