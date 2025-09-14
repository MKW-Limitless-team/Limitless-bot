CREATE TABLE
    IF NOT EXISTS playerdata (
        name TEXT,
        friend_code TEXT,
        discord_id TEXT PRIMARY KEY,
        mmr INTEGER,
        mii TEXT
    );

-- DROP TABLE playerdata