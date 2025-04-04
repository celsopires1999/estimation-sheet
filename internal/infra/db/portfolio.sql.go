// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: portfolio.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countPortfoliosByBaselineId = `-- name: CountPortfoliosByBaselineId :one
SELECT count(*) FROM portfolios WHERE baseline_id = $1
`

func (q *Queries) CountPortfoliosByBaselineId(ctx context.Context, baselineID string) (int64, error) {
	row := q.db.QueryRow(ctx, countPortfoliosByBaselineId, baselineID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countPortfoliosByPlanId = `-- name: CountPortfoliosByPlanId :one
SELECT count(*) FROM portfolios WHERE plan_id = $1
`

func (q *Queries) CountPortfoliosByPlanId(ctx context.Context, planID string) (int64, error) {
	row := q.db.QueryRow(ctx, countPortfoliosByPlanId, planID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deletePortfolio = `-- name: DeletePortfolio :one
DELETE FROM portfolios WHERE portfolio_id = $1 RETURNING portfolio_id, baseline_id, plan_id, start_date, created_at, updated_at
`

func (q *Queries) DeletePortfolio(ctx context.Context, portfolioID string) (Portfolio, error) {
	row := q.db.QueryRow(ctx, deletePortfolio, portfolioID)
	var i Portfolio
	err := row.Scan(
		&i.PortfolioID,
		&i.BaselineID,
		&i.PlanID,
		&i.StartDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findAllPortfoliosByBaselineIdWithRelations = `-- name: FindAllPortfoliosByBaselineIdWithRelations :many
SELECT
    pf.portfolio_id AS portfolio_id,
    bl.baseline_id AS baseline_id,
    pl.plan_type AS plan_type,
    pl.code AS plan_code,
    bl.code AS code,
    bl.review AS review,
    bl.title AS title,
    bl.description AS description,
    pf.start_date AS start_date,
    bl.duration AS duration,
    ma.name AS manager,
    es.name AS estimator,
    pf.created_at AS created_at,
    pf.updated_at AS updated_at
FROM
    baselines AS bl
    INNER JOIN portfolios AS pf ON bl.baseline_id = pf.baseline_id
    INNER JOIN users AS ma ON ma.user_id = bl.manager_id
    INNER JOIN users AS es ON es.user_id = bl.estimator_id
    INNER JOIN plans AS pl ON pl.plan_id = pf.plan_id
WHERE
    bl.baseline_id = $1
ORDER BY pl.code ASC
`

type FindAllPortfoliosByBaselineIdWithRelationsRow struct {
	PortfolioID string
	BaselineID  string
	PlanType    string
	PlanCode    string
	Code        string
	Review      int32
	Title       string
	Description pgtype.Text
	StartDate   pgtype.Date
	Duration    int32
	Manager     string
	Estimator   string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) FindAllPortfoliosByBaselineIdWithRelations(ctx context.Context, baselineID string) ([]FindAllPortfoliosByBaselineIdWithRelationsRow, error) {
	rows, err := q.db.Query(ctx, findAllPortfoliosByBaselineIdWithRelations, baselineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAllPortfoliosByBaselineIdWithRelationsRow
	for rows.Next() {
		var i FindAllPortfoliosByBaselineIdWithRelationsRow
		if err := rows.Scan(
			&i.PortfolioID,
			&i.BaselineID,
			&i.PlanType,
			&i.PlanCode,
			&i.Code,
			&i.Review,
			&i.Title,
			&i.Description,
			&i.StartDate,
			&i.Duration,
			&i.Manager,
			&i.Estimator,
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

const findAllPortfoliosByPlanIdWithRelations = `-- name: FindAllPortfoliosByPlanIdWithRelations :many
SELECT
    pf.portfolio_id AS portfolio_id,
    bl.baseline_id AS baseline_id,
    pl.plan_type AS plan_type,
    pl.code AS plan_code,
    bl.code AS code,
    bl.review AS review,
    bl.title AS title,
    bl.description AS description,
    pf.start_date AS start_date,
    bl.duration AS duration,
    ma.name AS manager,
    es.name AS estimator,
    pf.created_at AS created_at,
    pf.updated_at AS updated_at
FROM
    baselines AS bl
    INNER JOIN portfolios AS pf ON bl.baseline_id = pf.baseline_id
    INNER JOIN users AS ma ON ma.user_id = bl.manager_id
    INNER JOIN users AS es ON es.user_id = bl.estimator_id
    INNER JOIN plans AS pl ON pl.plan_id = pf.plan_id
WHERE
    pf.plan_id = $1
ORDER BY bl.code, pl.code ASC
`

type FindAllPortfoliosByPlanIdWithRelationsRow struct {
	PortfolioID string
	BaselineID  string
	PlanType    string
	PlanCode    string
	Code        string
	Review      int32
	Title       string
	Description pgtype.Text
	StartDate   pgtype.Date
	Duration    int32
	Manager     string
	Estimator   string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) FindAllPortfoliosByPlanIdWithRelations(ctx context.Context, planID string) ([]FindAllPortfoliosByPlanIdWithRelationsRow, error) {
	rows, err := q.db.Query(ctx, findAllPortfoliosByPlanIdWithRelations, planID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAllPortfoliosByPlanIdWithRelationsRow
	for rows.Next() {
		var i FindAllPortfoliosByPlanIdWithRelationsRow
		if err := rows.Scan(
			&i.PortfolioID,
			&i.BaselineID,
			&i.PlanType,
			&i.PlanCode,
			&i.Code,
			&i.Review,
			&i.Title,
			&i.Description,
			&i.StartDate,
			&i.Duration,
			&i.Manager,
			&i.Estimator,
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

const findAllPortfoliosWithRelations = `-- name: FindAllPortfoliosWithRelations :many
SELECT
    pf.portfolio_id AS portfolio_id,
    bl.baseline_id AS baseline_id,
    pl.plan_type AS plan_type,
    pl.code AS plan_code,
    bl.code AS code,
    bl.review AS review,
    bl.title AS title,
    bl.description AS description,
    pf.start_date AS start_date,
    bl.duration AS duration,
    ma.name AS manager,
    es.name AS estimator,
    pf.created_at AS created_at,
    pf.updated_at AS updated_at
FROM
    baselines AS bl
    INNER JOIN portfolios AS pf ON bl.baseline_id = pf.baseline_id
    INNER JOIN users AS ma ON ma.user_id = bl.manager_id
    INNER JOIN users AS es ON es.user_id = bl.estimator_id
    INNER JOIN plans AS pl ON pl.plan_id = pf.plan_id
ORDER BY bl.code ASC, bl.review DESC, pl.code ASC
`

type FindAllPortfoliosWithRelationsRow struct {
	PortfolioID string
	BaselineID  string
	PlanType    string
	PlanCode    string
	Code        string
	Review      int32
	Title       string
	Description pgtype.Text
	StartDate   pgtype.Date
	Duration    int32
	Manager     string
	Estimator   string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) FindAllPortfoliosWithRelations(ctx context.Context) ([]FindAllPortfoliosWithRelationsRow, error) {
	rows, err := q.db.Query(ctx, findAllPortfoliosWithRelations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAllPortfoliosWithRelationsRow
	for rows.Next() {
		var i FindAllPortfoliosWithRelationsRow
		if err := rows.Scan(
			&i.PortfolioID,
			&i.BaselineID,
			&i.PlanType,
			&i.PlanCode,
			&i.Code,
			&i.Review,
			&i.Title,
			&i.Description,
			&i.StartDate,
			&i.Duration,
			&i.Manager,
			&i.Estimator,
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

const findPortfolioById = `-- name: FindPortfolioById :one
SELECT portfolio_id, baseline_id, plan_id, start_date, created_at, updated_at FROM portfolios WHERE portfolio_id = $1
`

func (q *Queries) FindPortfolioById(ctx context.Context, portfolioID string) (Portfolio, error) {
	row := q.db.QueryRow(ctx, findPortfolioById, portfolioID)
	var i Portfolio
	err := row.Scan(
		&i.PortfolioID,
		&i.BaselineID,
		&i.PlanID,
		&i.StartDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPortfolioByIdWithRelations = `-- name: FindPortfolioByIdWithRelations :one
SELECT
    pf.portfolio_id AS portfolio_id,
    bl.baseline_id AS baseline_id,
    pl.plan_type AS plan_type,
    pl.code AS plan_code,
    bl.code AS code,
    bl.review AS review,
    bl.title AS title,
    bl.description AS description,
    pf.start_date AS start_date,
    bl.duration AS duration,
    ma.name AS manager,
    es.name AS estimator,
    pf.created_at AS created_at,
    pf.updated_at AS updated_at
FROM
    baselines AS bl
    INNER JOIN portfolios AS pf ON bl.baseline_id = pf.baseline_id
    INNER JOIN users AS ma ON ma.user_id = bl.manager_id
    INNER JOIN users AS es ON es.user_id = bl.estimator_id
    INNER JOIN plans AS pl ON pl.plan_id = pf.plan_id
WHERE
    pf.portfolio_id = $1
`

type FindPortfolioByIdWithRelationsRow struct {
	PortfolioID string
	BaselineID  string
	PlanType    string
	PlanCode    string
	Code        string
	Review      int32
	Title       string
	Description pgtype.Text
	StartDate   pgtype.Date
	Duration    int32
	Manager     string
	Estimator   string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) FindPortfolioByIdWithRelations(ctx context.Context, portfolioID string) (FindPortfolioByIdWithRelationsRow, error) {
	row := q.db.QueryRow(ctx, findPortfolioByIdWithRelations, portfolioID)
	var i FindPortfolioByIdWithRelationsRow
	err := row.Scan(
		&i.PortfolioID,
		&i.BaselineID,
		&i.PlanType,
		&i.PlanCode,
		&i.Code,
		&i.Review,
		&i.Title,
		&i.Description,
		&i.StartDate,
		&i.Duration,
		&i.Manager,
		&i.Estimator,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPortfolioByPlanIdAndBaselineCode = `-- name: FindPortfolioByPlanIdAndBaselineCode :one
SELECT portfolios.portfolio_id, portfolios.baseline_id, portfolios.plan_id, portfolios.start_date, portfolios.created_at, portfolios.updated_at
FROM portfolios
    INNER JOIN baselines ON baselines.baseline_id = portfolios.baseline_id
WHERE
    portfolios.plan_id = $1
    AND baselines.code = $2
LIMIT 1
`

type FindPortfolioByPlanIdAndBaselineCodeParams struct {
	PlanID string
	Code   string
}

func (q *Queries) FindPortfolioByPlanIdAndBaselineCode(ctx context.Context, arg FindPortfolioByPlanIdAndBaselineCodeParams) (Portfolio, error) {
	row := q.db.QueryRow(ctx, findPortfolioByPlanIdAndBaselineCode, arg.PlanID, arg.Code)
	var i Portfolio
	err := row.Scan(
		&i.PortfolioID,
		&i.BaselineID,
		&i.PlanID,
		&i.StartDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertPortfolio = `-- name: InsertPortfolio :exec
INSERT INTO
    portfolios (
        portfolio_id,
        baseline_id,
        plan_id,
        start_date,
        created_at
    )
VALUES ($1, $2, $3, $4, $5)
`

type InsertPortfolioParams struct {
	PortfolioID string
	BaselineID  string
	PlanID      string
	StartDate   pgtype.Date
	CreatedAt   pgtype.Timestamp
}

func (q *Queries) InsertPortfolio(ctx context.Context, arg InsertPortfolioParams) error {
	_, err := q.db.Exec(ctx, insertPortfolio,
		arg.PortfolioID,
		arg.BaselineID,
		arg.PlanID,
		arg.StartDate,
		arg.CreatedAt,
	)
	return err
}

const updatePortfolio = `-- name: UpdatePortfolio :exec
UPDATE portfolios
SET
    baseline_id = $2,
    plan_id = $3,
    start_date = $4,
    updated_at = $5
WHERE
    portfolio_id = $1
`

type UpdatePortfolioParams struct {
	PortfolioID string
	BaselineID  string
	PlanID      string
	StartDate   pgtype.Date
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) UpdatePortfolio(ctx context.Context, arg UpdatePortfolioParams) error {
	_, err := q.db.Exec(ctx, updatePortfolio,
		arg.PortfolioID,
		arg.BaselineID,
		arg.PlanID,
		arg.StartDate,
		arg.UpdatedAt,
	)
	return err
}
