CREATE TABLE purchases
(
    id          SERIAL PRIMARY KEY,
    "name"      VARCHAR(255),
    description TEXT,
    "count"     INTEGER,
    unit        VARCHAR(255),
    price       INTEGER,
    created_at  TIMESTAMP,
    owner_id    INTEGER,
    "groups"    INTEGER ARRAY
);

CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    "name"      VARCHAR(255),
    total_count INTEGER,
    owner_id    INTEGER,
    "groups"    INTEGER ARRAY
);

CREATE TABLE recipes
(
    id          SERIAL PRIMARY KEY,
    "name"      VARCHAR(255),
    description TEXT,
    ingredients TEXT,
    actions     TEXT,
    Complexity  INTEGER,
    owner_id    INTEGER,
    "groups"    INTEGER ARRAY
);

CREATE TABLE workouts
(
    id           SERIAL PRIMARY KEY,
    payment_date TIMESTAMP,
    end_date     TIMESTAMP,
    Remains      INTEGER,
    owner_id     INTEGER,
    "groups"     INTEGER ARRAY
);