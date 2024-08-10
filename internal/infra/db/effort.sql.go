// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: effort.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const bulkInsertEffort = `-- name: BulkInsertEffort :exec
INSERT INTO
    efforts (
        effort_id,
        baseline_id,
        competence_id,
        comment,
        hours,
        created_at
    )
VALUES (
        unnest($1::text []),
        unnest($2::text []),
        unnest($3::text []),
        unnest($4::text []),
        unnest($5::int[]),
        unnest($6::timestamp[])
    )
`

type BulkInsertEffortParams struct {
	Column1 []string
	Column2 []string
	Column3 []string
	Column4 []string
	Column5 []int32
	Column6 []pgtype.Timestamp
}

func (q *Queries) BulkInsertEffort(ctx context.Context, arg BulkInsertEffortParams) error {
	_, err := q.db.Exec(ctx, bulkInsertEffort,
		arg.Column1,
		arg.Column2,
		arg.Column3,
		arg.Column4,
		arg.Column5,
		arg.Column6,
	)
	return err
}

const bulkInsertEffortAllocation = `-- name: BulkInsertEffortAllocation :exec
INSERT INTO
    effort_allocations (
        effort_allocation_id,
        effort_id,
        allocation_date,
        hours,
        created_at
    )
VALUES (
        unnest($1::text []),
        unnest($2::text []),
        unnest($3::date[]),
        unnest($4::int[]),
        unnest($5::timestamp[])
    )
`

type BulkInsertEffortAllocationParams struct {
	Column1 []string
	Column2 []string
	Column3 []pgtype.Date
	Column4 []int32
	Column5 []pgtype.Timestamp
}

func (q *Queries) BulkInsertEffortAllocation(ctx context.Context, arg BulkInsertEffortAllocationParams) error {
	_, err := q.db.Exec(ctx, bulkInsertEffortAllocation,
		arg.Column1,
		arg.Column2,
		arg.Column3,
		arg.Column4,
		arg.Column5,
	)
	return err
}

const deleteEffort = `-- name: DeleteEffort :execrows
DELETE FROM efforts WHERE effort_id = $1
`

func (q *Queries) DeleteEffort(ctx context.Context, effortID string) (int64, error) {
	result, err := q.db.Exec(ctx, deleteEffort, effortID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const deleteEffortAllocations = `-- name: DeleteEffortAllocations :execrows
DELETE FROM effort_allocations WHERE effort_id = $1
`

func (q *Queries) DeleteEffortAllocations(ctx context.Context, effortID string) (int64, error) {
	result, err := q.db.Exec(ctx, deleteEffortAllocations, effortID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findEffortAllocationsByEffortId = `-- name: FindEffortAllocationsByEffortId :many
SELECT effort_allocation_id, effort_id, allocation_date, hours, created_at, updated_at
FROM effort_allocations
WHERE
    effort_id = $1
ORDER BY allocation_date ASC
`

func (q *Queries) FindEffortAllocationsByEffortId(ctx context.Context, effortID string) ([]EffortAllocation, error) {
	rows, err := q.db.Query(ctx, findEffortAllocationsByEffortId, effortID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EffortAllocation
	for rows.Next() {
		var i EffortAllocation
		if err := rows.Scan(
			&i.EffortAllocationID,
			&i.EffortID,
			&i.AllocationDate,
			&i.Hours,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const findEffortById = `-- name: FindEffortById :one
SELECT effort_id, baseline_id, competence_id, comment, hours, created_at, updated_at FROM efforts WHERE effort_id = $1
`

func (q *Queries) FindEffortById(ctx context.Context, effortID string) (Effort, error) {
	row := q.db.QueryRow(ctx, findEffortById, effortID)
	var i Effort
	err := row.Scan(
		&i.EffortID,
		&i.BaselineID,
		&i.CompetenceID,
		&i.Comment,
		&i.Hours,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findEffortsByBaselineIdWithRelations = `-- name: FindEffortsByBaselineIdWithRelations :many
SELECT
    e.effort_id AS effort_id,
    e.baseline_id AS baseline_id,
    e.competence_id AS competence_id,
    c.code AS competence_code,
    c.name AS competence_name,
    e.comment AS comment,
    e.hours AS hours,
    e.created_at AS created_at,
    e.updated_at AS updated_at
FROM efforts AS e
    INNER JOIN competences AS c ON e.competence_id = c.competence_id
WHERE
    e.baseline_id = $1
ORDER BY c.code ASC
`

type FindEffortsByBaselineIdWithRelationsRow struct {
	EffortID       string
	BaselineID     string
	CompetenceID   string
	CompetenceCode string
	CompetenceName string
	Comment        pgtype.Text
	Hours          int32
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

func (q *Queries) FindEffortsByBaselineIdWithRelations(ctx context.Context, baselineID string) ([]FindEffortsByBaselineIdWithRelationsRow, error) {
	rows, err := q.db.Query(ctx, findEffortsByBaselineIdWithRelations, baselineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindEffortsByBaselineIdWithRelationsRow
	for rows.Next() {
		var i FindEffortsByBaselineIdWithRelationsRow
		if err := rows.Scan(
			&i.EffortID,
			&i.BaselineID,
			&i.CompetenceID,
			&i.CompetenceCode,
			&i.CompetenceName,
			&i.Comment,
			&i.Hours,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const insertEffort = `-- name: InsertEffort :exec
INSERT INTO
    efforts (
        effort_id,
        baseline_id,
        competence_id,
        comment,
        hours,
        created_at
    )
VALUES ($1, $2, $3, $4, $5, $6)
`

type InsertEffortParams struct {
	EffortID     string
	BaselineID   string
	CompetenceID string
	Comment      pgtype.Text
	Hours        int32
	CreatedAt    pgtype.Timestamp
}

func (q *Queries) InsertEffort(ctx context.Context, arg InsertEffortParams) error {
	_, err := q.db.Exec(ctx, insertEffort,
		arg.EffortID,
		arg.BaselineID,
		arg.CompetenceID,
		arg.Comment,
		arg.Hours,
		arg.CreatedAt,
	)
	return err
}

const insertEffortAllocation = `-- name: InsertEffortAllocation :exec
INSERT INTO
    effort_allocations (
        effort_allocation_id,
        effort_id,
        allocation_date,
        hours,
        created_at
    )
VALUES ($1, $2, $3, $4, $5)
`

type InsertEffortAllocationParams struct {
	EffortAllocationID string
	EffortID           string
	AllocationDate     pgtype.Date
	Hours              int32
	CreatedAt          pgtype.Timestamp
}

func (q *Queries) InsertEffortAllocation(ctx context.Context, arg InsertEffortAllocationParams) error {
	_, err := q.db.Exec(ctx, insertEffortAllocation,
		arg.EffortAllocationID,
		arg.EffortID,
		arg.AllocationDate,
		arg.Hours,
		arg.CreatedAt,
	)
	return err
}

const updateEffort = `-- name: UpdateEffort :exec
UPDATE efforts
SET
    baseline_id = $2,
    competence_id = $3,
    comment = $4,
    hours = $5,
    updated_at = $6
WHERE
    effort_id = $1
`

type UpdateEffortParams struct {
	EffortID     string
	BaselineID   string
	CompetenceID string
	Comment      pgtype.Text
	Hours        int32
	UpdatedAt    pgtype.Timestamp
}

func (q *Queries) UpdateEffort(ctx context.Context, arg UpdateEffortParams) error {
	_, err := q.db.Exec(ctx, updateEffort,
		arg.EffortID,
		arg.BaselineID,
		arg.CompetenceID,
		arg.Comment,
		arg.Hours,
		arg.UpdatedAt,
	)
	return err
}
