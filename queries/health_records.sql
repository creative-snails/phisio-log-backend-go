-- name: CreateHealthRecord :one
INSERT INTO health_records (
    -- user_id,
    parent_record_id,
    description,
    progress,
    improvement,
    severity,
    treatments_tried
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetHealthRecord :one
SELECT * FROM health_records
WHERE id = $1 LIMIT 1;

-- name: ListHealthRecords :many
-- SELECT * FROM health_records
-- WHERE user_id = $1
-- ORDER BY created_at DESC;