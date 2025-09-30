CREATE TABLE
    IF NOT EXISTS placements (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        track TEXT,
        discord_id TEXT,
        minutes INTEGER,
        seconds INTEGER,
        milliseconds INTEGER,
        character TEXT,
        vehicle TEXT,
        drift_type TEXT,
        category TEXT,
        url TEXT,
        crc INTEGER UNIQUE,
        approved BOOLEAN,
        FOREIGN KEY (discord_id) REFERENCES playerdata (discord_id)
    );

-- DROP TABLE placements