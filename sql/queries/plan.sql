-- name: InsertPlan :exec
INSERT INTO
    plans (
        plan_id,
        plan_type,
        code,
        name,
        assumptions,
        created_at
    )
VALUES ($1, $2, $3, $4, $5, $6);

-- name: FindPlanById :one
SELECT * FROM plans WHERE plan_id = $1;

-- name: FindPlanByCode :one
SELECT * FROM plans WHERE code = $1;

-- name: FindAllPlans :many
SELECT * FROM plans ORDER BY code ASC;

-- name: UpdatePlan :one
UPDATE plans
SET
    plan_type = $2,
    code = $3,
    name = $4,
    assumptions = $5,
    updated_at = $6
WHERE
    plan_id = $1
RETURNING
    *;

-- name: DeletePlan :one
DELETE FROM plans WHERE plan_id = $1 RETURNING *;