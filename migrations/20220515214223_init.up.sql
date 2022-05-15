CREATE TABLE purchases
(
    id          SERIAL PRIMARY KEY,
    "name"      VARCHAR(255),
    description TEXT,
    "count"     INTEGER,
    unit        VARCHAR(4),
    price       INTEGER,
    created_at  TIMESTAMP
);

CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    "name"      VARCHAR(255),
    total_count INTEGER,
);

CREATE TABLE recipes
(
    id          SERIAL PRIMARY KEY,
    "name"      VARCHAR(255),
    description TEXT,
    ingredients TEXT,
    actions     TEXT,
    Complexity  INTEGER,
);

CREATE TABLE workouts
(
    id           SERIAL PRIMARY KEY,
    payment_date TIMESTAMP,
    end_date     TIMESTAMP,
    Remains      INTEGER,
);