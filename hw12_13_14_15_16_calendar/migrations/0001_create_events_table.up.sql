-- 0001_create_events_table.up.sql
CREATE TABLE events
(
    id         SERIAL PRIMARY KEY,
    title      TEXT      NOT NULL,
    start_time TIMESTAMP NOT NULL
);