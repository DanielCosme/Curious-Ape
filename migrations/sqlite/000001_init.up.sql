-- schema
CREATE TABLE IF NOT EXISTS days (
        id                  INTEGER PRIMARY KEY,
        "date"              DATE NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS days_date_idx ON days ("date");

CREATE TABLE IF NOT EXISTS habit_categories (
        id                  INTEGER PRIMARY KEY,

        name                TEXT NOT NULL                           CHECK (LENGTH(name) > 0),
        type                TEXT UNIQUE NOT NULL DEFAULT "dynamic"  CHECK (LENGTH(type) > 0),
        description         TEXT DEFAULT ""

        CHECK(LENGTH(id) > 0)
);

CREATE TABLE IF NOT EXISTS habits (
        id                  INTEGER PRIMARY KEY,
        day_id              INTEGER NOT NULL,
        habit_category_id   INTEGER NOT NULL,

        status              TEXT CHECK (status IN ('no_info', 'not_done', 'done')),

        FOREIGN KEY (habit_category_id) REFERENCES habit_categories (id) ON DELETE CASCADE,
        FOREIGN KEY (day_id) REFERENCES "days" (id) ON DELETE CASCADE,
        UNIQUE (day_id, habit_category_id),
        CHECK(length(id) > 0 AND length(day_id) > 0 AND length(habit_category_id) > 0)
);

CREATE TABLE IF NOT EXISTS habit_logs (
        id                  INTEGER PRIMARY KEY,
        habit_id            INTEGER NOT NULL,

        origin              TEXT NOT NULL DEFAULT "unknown" CHECK (length(origin) > 0),
        success             BOOLEAN DEFAULT false,
        is_automated        BOOLEAN NOT NULL DEFAULT false,
        note                TEXT DEFAULT "",

        FOREIGN KEY (habit_id) REFERENCES habits (id) ON DELETE CASCADE,
        UNIQUE (habit_id, origin),
        CHECK(length(id) > 0 AND length(habit_id) > 0)
);

CREATE TABLE IF NOT EXISTS auths (
        id                  INTEGER primary key,

        provider            TEXT not null UNIQUE,
        access_token        TEXT not null UNIQUE,
        refresh_token       TEXT,
        token_type          TEXT,
        expiration          DATE
);

CREATE TABLE IF NOT EXISTS users (
        id              INTEGER primary key,

        name            TEXT not null UNIQUE,
        password        TEXT not null UNIQUE,
        role            TEXT CHECK (role IN ('admin', 'user', 'guest')) NOT NULL,
        email           TEXT NOT NULL default ''
);

CREATE TABLE IF NOT EXISTS sleep_logs (
        id                  INTEGER primary key,
        day_id              INTEGER not null,

        "date"              DATE not null,
        start_time          DATE not null,
        end_time            DATE not null,
        is_main_sleep       BOOLEAN default true,
        is_automated        BOOLEAN default false,
        origin              TEXT not null CHECK (length(origin) > 1),
        total_time_in_bed   INTEGER default 0,
        minutes_asleep      INTEGER default 0,
        minutes_rem         INTEGER default 0,
        minutes_deep        INTEGER default 0,
        minutes_light       INTEGER default 0,
        minutes_awake       INTEGER default 0,
        raw                 TEXT,

        FOREIGN KEY (day_id) REFERENCES "days" (id) ON DELETE CASCADE,
        UNIQUE (day_id, is_main_sleep)
);

CREATE TABLE IF NOT EXISTS fitness_logs (
        id                  INTEGER primary key,
        day_id              INTEGER not null,

        "date"              DATE not null,
        start_time          DATE not null,
        end_time            DATE not null,
        "type"              TEXT not null default '',
        title               TEXT not null default '',
        origin              TEXT not null CHECK (length(origin) > 1),
        note                TEXT default '',
        raw                 TEXT,

        FOREIGN KEY (day_id) REFERENCES "days" (id) ON DELETE CASCADE,
        UNIQUE (day_id, start_time)
);

CREATE TABLE IF NOT EXISTS sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry);
