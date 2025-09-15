CREATE TABLE
    IF NOT EXISTS placements (
        track TEXT,
        discord_id TEXT PRIMARY KEY,
        flag TEXT,
        time TEXT,
        character TEXT,
        vehicle TEXT,
        drift_type TEXT,
        category TEXT,
        approved BOOLEAN
    );