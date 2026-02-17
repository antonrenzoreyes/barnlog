CREATE TABLE events (
    id TEXT PRIMARY KEY,
    aggregate_type TEXT NOT NULL CHECK (length(trim(aggregate_type)) > 0),
    aggregate_id TEXT NOT NULL CHECK (length(trim(aggregate_id)) > 0),
    event_type TEXT NOT NULL CHECK (length(trim(event_type)) > 0),
    created_by TEXT NOT NULL CHECK (length(trim(created_by)) > 0),
    source TEXT NOT NULL CHECK (length(trim(source)) > 0),
    request_id TEXT NOT NULL CHECK (length(trim(request_id)) > 0),
    event_version INTEGER NOT NULL DEFAULT 1,
    payload_json TEXT NOT NULL,
    metadata_json TEXT,
    occurred_at TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (datetime('now'))
);
CREATE INDEX idx_events_aggregate
    ON events (aggregate_type, aggregate_id, occurred_at);
CREATE INDEX idx_events_type_time
    ON events (event_type, occurred_at);
CREATE UNIQUE INDEX ux_events_source_request_id
    ON events (source, request_id);
CREATE UNIQUE INDEX version_unique ON schema_migrations (version);
