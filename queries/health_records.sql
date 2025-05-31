-- =============================================
-- Health Records
-- =============================================

-- name: CreateHealthRecord :one
INSERT INTO health_records (
    parent_record_id,
    description,
    stage,
    severity,
    progression,
    treatments_tried
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetHealthRecord :one
SELECT * FROM health_records
WHERE id = $1 LIMIT 1;

-- =============================================
-- Symptoms
-- =============================================

-- name: CreateSymptom :one
INSERT INTO symptoms (
    health_record_id,
    name,
    start_date
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetSymptoms :many
SELECT * FROM symptoms
WHERE health_record_id = $1;

-- name: UpdateSymptom :one
UPDATE symptoms
SET 
    name = $2,
    start_date = $3
WHERE id = $1
RETURNING *;

-- name: DeleteSymptom :one
DELETE FROM symptoms
WHERE id = $1
RETURNING *;

-- =============================================
-- Affected Parts
-- =============================================

-- name: CreateAffectedPart :one
INSERT INTO affected_parts (
    symptom_id,
    body_part_id,
    state
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAffectedParts :many
SELECT * FROM affected_parts
WHERE symptom_id = $1;

-- name: UpdateAffectedPart :one
UPDATE affected_parts
SET
    state = $3
WHERE (symptom_id = $1 AND body_part_id = $2)
RETURNING *;

-- name: DeleteAffectedPart :one
DELETE FROM affected_parts
WHERE (symptom_id = $1 AND body_part_id = $2)
RETURNING *;