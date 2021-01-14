CREATE TABLE resource_type (
    -- Fields
    resource_type_id INTEGER PRIMARY KEY AUTOINCREMENT,
    alias            TEXT    NOT NULL,

    -- Constraints
    UNIQUE(alias)
);

INSERT INTO resource_type (alias)
VALUES ('ping'), ('http');

CREATE TABLE resource (
    -- Fields
    resource_id      TEXT    NOT NULL PRIMARY KEY,
    alias            TEXT    NOT NULL,
    active           BOOLEAN NOT NULL DEFAULT true,
    resource_type_id INTEGER NOT NULL,
    check_interval   INTEGER NOT NULL,
    next_check       INTEGER NOT NULL,

    -- Constraints
    CHECK(check_interval > 0),
    CHECK(next_check > -1),
    FOREIGN KEY (resource_type_id)
        REFERENCES resource_type(resource_type_id)
        ON DELETE CASCADE
);

CREATE TABLE resource_status (
    -- Fields
    resource_status_id INTEGER PRIMARY KEY AUTOINCREMENT,
    resource_id        TEXT    NOT NULL,
    status             TEXT    NOT NULL,
    details            TEXT,
    timestamp          INTEGER NOT NULL, -- UNIX timestamp as a 64 bit integer.

    -- Constraints
    CHECK (status IN ('OK', 'WARN', 'ALARM')),
    FOREIGN KEY (resource_id)
        REFERENCES resource(resource_id)
        ON DELETE CASCADE
)
