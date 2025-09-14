CREATE TABLE
    IF NOT EXISTS characters (name TEXT PRIMARY KEY, size TEXT);

INSERT INTO
    characters (name, size)
VALUES
    ('Baby Mario', 'S'),
    ('Baby Luigi', 'S'),
    ('Baby Peach', 'S'),
    ('Baby Daisy', 'S'),
    ('Toad', 'S'),
    ('Toadette', 'S'),
    ('Koopa Troopa', 'S'),
    ('Dry Bones', 'S'),
    ('Mario', 'M'),
    ('Luigi', 'M'),
    ('Peach', 'M'),
    ('Daisy', 'M'),
    ('Yoshi', 'M'),
    ('Birdo', 'M'),
    ('Diddy Kong', 'M'),
    ('Bowser Jr', 'M'),
    ('Wario', 'L'),
    ('Waluigi', 'L'),
    ('Donkey Kong', 'L'),
    ('Funky Kong', 'L'),
    ('Bowser', 'L'),
    ('King Boo', 'L'),
    ('Rosalina', 'L'),
    ('Dry Bowser', 'L'),
    ('Mii', '');

-- DROP TABLE characters