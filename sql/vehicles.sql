CREATE TABLE
    IF NOT EXISTS vehicles (id INTEGER PRIMARY KEY, name TEXT, size TEXT);

INSERT INTO
    vehicles (id, name, size)
VALUES
    (0, 'Standard Kart S', 'S'),
    (3, 'Booster Seat', 'S'),
    (6, 'Mini Beast', 'S'),
    (9, 'Tiny Titan', 'S'),
    (12, 'Blue Falcon', 'S'),
    (15, 'Cheep Charger', 'S'),
    (18, 'Standard Bike S', 'S'),
    (21, 'Bullet Bike', 'S'),
    (24, 'Bit Bike', 'S'),
    (27, 'Quacker', 'S'),
    (30, 'Magikruiser', 'S'),
    (33, 'Jet Bubble', 'S'),
    (1, 'Standard Kart M', 'M'),
    (4, 'Classic Dragster', 'M'),
    (7, 'Wild Wing', 'M'),
    (10, 'Super Blooper', 'M'),
    (13, 'Daytripper', 'M'),
    (16, 'Sprinter', 'M'),
    (19, 'Standard Bike M', 'M'),
    (22, 'Mach Bike', 'M'),
    (25, 'Sugarscoot', 'M'),
    (28, 'Zip Zip', 'M'),
    (31, 'Sneakster', 'M'),
    (34, 'Dolphin Dasher', 'M'),
    (2, 'Standard Kart L', 'L'),
    (5, 'Offroader', 'L'),
    (8, 'Flame Flyer', 'L'),
    (11, 'Piranha Prowler', 'L'),
    (14, 'Jetsetter', 'L'),
    (17, 'Honeycoupe', 'L'),
    (20, 'Standard Bike L', 'L'),
    (23, 'Flame Runner', 'L'),
    (26, 'Wario Bike', 'L'),
    (29, 'Shooting Star', 'L'),
    (32, 'Spear', 'L'),
    (35, 'Phantom', 'L');

-- DROP TABLE vehicles