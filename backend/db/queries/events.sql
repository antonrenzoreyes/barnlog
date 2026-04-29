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

-- name: GetEventBySourceRequestID :one
SELECT
    id,
    aggregate_type,
    aggregate_id,
    event_type,
    payload_json
FROM events
WHERE source = ? AND request_id = ?
LIMIT 1;
