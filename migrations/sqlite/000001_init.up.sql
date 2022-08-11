-- schema
CREATE TABLE IF NOT EXISTS days (
        id                  INTEGER primary key,
        "date"              DATE not null UNIQUE,
        deep_work_minutes   INTEGER NOT NULL DEFAULT 0 check ( length(deep_work_minutes) >= 0)
);

CREATE TABLE IF NOT EXISTS habit_categories (
        id                  INTEGER primary key,

        name                TEXT not null                   check (length(name) > 0),
        type                TEXT not null                   check (length(type) > 0),
        code                TEXT not null default "default" check (length(code) > 0),
        description         TEXT default "",
        color               INTEGER default "#ffffff",

        CHECK(length(id) > 0)
);

INSERT INTO habit_categories (name, type)
VALUES  ("Eat healthy", "food"),
        ("Wake up early", "wake_up"),
        ("Workout", "fitness"),
        ("Deep work", "deep_work");

CREATE TABLE IF NOT EXISTS habits (
        id                  INTEGER primary key,
        day_id              INTEGER not null,
        habit_category_id   INTEGER not null,

        status              TEXT not null check (length(status) > 0),

        FOREIGN KEY (habit_category_id) REFERENCES habit_categories (id) ON DELETE CASCADE,
        FOREIGN KEY (day_id) REFERENCES "days" (id) ON DELETE CASCADE,
        UNIQUE (day_id, habit_category_id),
        CHECK(length(id) > 0 AND length(day_id) > 0 AND length(habit_category_id) > 0)
);

CREATE TABLE IF NOT EXISTS habit_logs (
        id                  INTEGER primary key,
        habit_id            INTEGER not null,

        origin              TEXT not null default "unknown" check (length(origin) > 0),
        success             BOOLEAN default false,
        is_automated        BOOLEAN not null default false,
        note                TEXT default "",

        FOREIGN KEY (habit_id) REFERENCES habits (id) ON DELETE CASCADE,
        UNIQUE (habit_id, origin),
        CHECK(length(id) > 0 AND length(habit_id) > 0)
);


CREATE TABLE IF NOT EXISTS oauths (
        id                  INTEGER primary key,

        provider            TEXT not null UNIQUE,
        access_token        TEXT not null UNIQUE,
        refresh_token       TEXT,
        type                TEXT,
        expiration          DATE,

        toggl_workspace_id      INTEGER default "",
        toggl_organization_id   INTEGER default "",
        toggl_project_ids       TEXT defatult ""
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


