CREATE TABLE
    IF NOT EXISTS vehicles (name TEXT PRIMARY KEY, size TEXT);

INSERT INTO
    vehicles (name, size)
VALUES
    ('Standard Kart S', 'S'),
    ('Booster Seat', 'S'),
    ('Mini Beast', 'S'),
    ('Tiny Titan', 'S'),
    ('Blue Falcon', 'S'),
    ('Cheep Charger', 'S'),
    ('Standard Bike S', 'S'),
    ('Bullet Bike', 'S'),
    ('Bit Bike', 'S'),
    ('Quacker', 'S'),
    ('Magikruiser', 'S'),
    ('Jet Bubble', 'S'),
    ('Standard Kart M', 'M'),
    ('Classic Dragster', 'M'),
    ('Wild Wing', 'M'),
    ('Super Blooper', 'M'),
    ('Daytripper', 'M'),
    ('Sprinter', 'M'),
    ('Standard Bike M', 'M'),
    ('Mach Bike', 'M'),
    ('Sugarscoot', 'M'),
    ('Zip Zip', 'M'),
    ('Sneakster', 'M'),
    ('Dolphin Dasher', 'M'),
    ('Standard Kart L', 'L'),
    ('Offroader', 'L'),
    ('Flame Flyer', 'L'),
    ('Piranha Prowler', 'L'),
    ('Jetsetter', 'L'),
    ('Honeycoupe', 'L'),
    ('Standard Bike L', 'L'),
    ('Flame Runner', 'L'),
    ('Wario Bike', 'L'),
    ('Shooting Star', 'L'),
    ('Spear', 'L'),
    ('Phantom', 'L');

-- DROP TABLE vehicles