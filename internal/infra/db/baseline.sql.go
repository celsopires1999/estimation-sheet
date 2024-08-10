// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: baseline.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteBaseline = `-- name: DeleteBaseline :one
DELETE FROM baselines WHERE baseline_id = $1 RETURNING baseline_id, code, review, title, description, start_date, duration, manager_id, estimator_id, created_at, updated_at
`

func (q *Queries) DeleteBaseline(ctx context.Context, baselineID string) (Baseline, error) {
	row := q.db.QueryRow(ctx, deleteBaseline, baselineID)
	var i Baseline
	err := row.Scan(
		&i.BaselineID,
		&i.Code,
		&i.Review,
		&i.Title,
		&i.Description,
		&i.StartDate,
		&i.Duration,
		&i.ManagerID,
		&i.EstimatorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findAllBaselines = `-- name: FindAllBaselines :many
SELECT baselines.baseline_id, baselines.code, baselines.review, baselines.title, baselines.description, baselines.start_date, baselines.duration, baselines.manager_id, baselines.estimator_id, baselines.created_at, baselines.updated_at, managers.name AS manager, estimators.name AS estimator
FROM
    baselines
    INNER JOIN users AS managers ON managers.user_id = baselines.manager_id
    INNER JOIN users AS estimators ON estimators.user_id = baselines.estimator_id
ORDER BY code ASC, review DESC
`

type FindAllBaselinesRow struct {
	BaselineID  string
	Code        string
	Review      int32
	Title       string
	Description pgtype.Text
	StartDate   pgtype.Date
	Duration    int32
	ManagerID   string
	EstimatorID string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	Manager     string
	Estimator   string
}

func (q *Queries) FindAllBaselines(ctx context.Context) ([]FindAllBaselinesRow, error) {
	rows, err := q.db.Query(ctx, findAllBaselines)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAllBaselinesRow
	for rows.Next() {
		var i FindAllBaselinesRow
		if err := rows.Scan(
			&i.BaselineID,
			&i.Code,
			&i.Review,
			&i.Title,
			&i.Description,
			&i.StartDate,
			&i.Duration,
			&i.ManagerID,
			&i.EstimatorID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Manager,
			&i.Estimator,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findBaselineById = `-- name: FindBaselineById :one
SELECT baseline_id, code, review, title, description, start_date, duration, manager_id, estimator_id, created_at, updated_at FROM baselines WHERE baseline_id = $1
`

func (q *Queries) FindBaselineById(ctx context.Context, baselineID string) (Baseline, error) {
	row := q.db.QueryRow(ctx, findBaselineById, baselineID)
	var i Baseline
	err := row.Scan(
		&i.BaselineID,
		&i.Code,
		&i.Review,
		&i.Title,
		&i.Description,
		&i.StartDate,
		&i.Duration,
		&i.ManagerID,
		&i.EstimatorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findBaselineByIdWithRelations = `-- name: FindBaselineByIdWithRelations :one
SELECT baselines.baseline_id, baselines.code, baselines.review, baselines.title, baselines.description, baselines.start_date, baselines.duration, baselines.manager_id, baselines.estimator_id, baselines.created_at, baselines.updated_at, managers.name AS manager, estimators.name AS estimator
FROM
    baselines
    INNER JOIN users AS managers ON managers.user_id = baselines.manager_id
    INNER JOIN users AS estimators ON estimators.user_id = baselines.estimator_id
WHERE
    baseline_id = $1
`

type FindBaselineByIdWithRelationsRow struct {
	BaselineID  string
	Code        string
	Review      int32
	Title       string
	Description pgtype.Text
	StartDate   pgtype.Date
	Duration    int32
	ManagerID   string
	EstimatorID string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	Manager     string
	Estimator   string
}

func (q *Queries) FindBaselineByIdWithRelations(ctx context.Context, baselineID string) (FindBaselineByIdWithRelationsRow, error) {
	row := q.db.QueryRow(ctx, findBaselineByIdWithRelations, baselineID)
	var i FindBaselineByIdWithRelationsRow
	err := row.Scan(
		&i.BaselineID,
		&i.Code,
		&i.Review,
		&i.Title,
		&i.Description,
		&i.StartDate,
		&i.Duration,
		&i.ManagerID,
		&i.EstimatorID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Manager,
		&i.Estimator,
	)
	return i, err
}

const insertBaseline = `-- name: InsertBaseline :exec
INSERT INTO
    baselines (
        baseline_id,
        code,
        review,
        title,
        description,
        start_date,
        duration,
        manager_id,
        estimator_id,
        created_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10
    )
`

type InsertBaselineParams struct {
	BaselineID  string
	Code        string
	Review      int32
	Title       string
	Description pgtype.Text
	StartDate   pgtype.Date
	Duration    int32
	ManagerID   string
	EstimatorID string
	CreatedAt   pgtype.Timestamp
}

func (q *Queries) InsertBaseline(ctx context.Context, arg InsertBaselineParams) error {
	_, err := q.db.Exec(ctx, insertBaseline,
		arg.BaselineID,
		arg.Code,
		arg.Review,
		arg.Title,
		arg.Description,
		arg.StartDate,
		arg.Duration,
		arg.ManagerID,
		arg.EstimatorID,
		arg.CreatedAt,
	)
	return err
}

const updateBaseline = `-- name: UpdateBaseline :exec
UPDATE baselines
SET
    code = $2,
    review = $3,
    title = $4,
    description = $5,
    start_date = $6,
    duration = $7,
    manager_id = $8,
    estimator_id = $9,
    updated_at = $10
WHERE
    baseline_id = $1
`

type UpdateBaselineParams struct {
	BaselineID  string
	Code        string
	Review      int32
	Title       string
	Description pgtype.Text
	StartDate   pgtype.Date
	Duration    int32
	ManagerID   string
	EstimatorID string
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) UpdateBaseline(ctx context.Context, arg UpdateBaselineParams) error {
	_, err := q.db.Exec(ctx, updateBaseline,
		arg.BaselineID,
		arg.Code,
		arg.Review,
		arg.Title,
		arg.Description,
		arg.StartDate,
		arg.Duration,
		arg.ManagerID,
		arg.EstimatorID,
		arg.UpdatedAt,
	)
	return err
}
