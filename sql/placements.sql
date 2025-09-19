CREATE TABLE
    IF NOT EXISTS placements (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        track TEXT,
        discord_id TEXT,
        flag TEXT,
        minutes INTEGER,
        seconds INTEGER,
        milliseconds INTEGER,
        character TEXT,
        vehicle TEXT,
        drift_type TEXT,
        category TEXT,
        url TEXT,
        approved BOOLEAN
    );

-- DROP TABLE placements