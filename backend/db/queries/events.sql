-- name: CreateEvent :exec
INSERT INTO events (
    id,
    aggregate_type,
    aggregate_id,
    event_type,
    created_by,
    source,
    request_id,
    event_version,
    payload_json,
    metadata_json,
    occurred_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);