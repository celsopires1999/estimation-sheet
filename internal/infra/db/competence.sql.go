// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: competence.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteCompetence = `-- name: DeleteCompetence :execrows
DELETE FROM competences WHERE competence_id = $1
`

func (q *Queries) DeleteCompetence(ctx context.Context, competenceID string) (int64, error) {
	result, err := q.db.Exec(ctx, deleteCompetence, competenceID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findAllCompetences = `-- name: FindAllCompetences :many
SELECT competence_id, code, name, created_at, updated_at FROM competences ORDER BY code ASC
`

func (q *Queries) FindAllCompetences(ctx context.Context) ([]Competence, error) {
	rows, err := q.db.Query(ctx, findAllCompetences)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Competence
	for rows.Next() {
		var i Competence
		if err := rows.Scan(
			&i.CompetenceID,
			&i.Code,
			&i.Name,
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

const findCompetenceById = `-- name: FindCompetenceById :one
SELECT competence_id, code, name, created_at, updated_at FROM competences WHERE competence_id = $1
`

func (q *Queries) FindCompetenceById(ctx context.Context, competenceID string) (Competence, error) {
	row := q.db.QueryRow(ctx, findCompetenceById, competenceID)
	var i Competence
	err := row.Scan(
		&i.CompetenceID,
		&i.Code,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertCompetence = `-- name: InsertCompetence :exec
INSERT INTO
    competences (
        competence_id,
        code,
        name,
        created_at
    )
VALUES ($1, $2, $3, $4)
`

type InsertCompetenceParams struct {
	CompetenceID string
	Code         string
	Name         string
	CreatedAt    pgtype.Timestamp
}

func (q *Queries) InsertCompetence(ctx context.Context, arg InsertCompetenceParams) error {
	_, err := q.db.Exec(ctx, insertCompetence,
		arg.CompetenceID,
		arg.Code,
		arg.Name,
		arg.CreatedAt,
	)
	return err
}

const updateCompetence = `-- name: UpdateCompetence :exec
UPDATE competences
SET
    code = $2,
    name = $3,
    updated_at = $4
WHERE
    competence_id = $1
RETURNING
    competence_id, code, name, created_at, updated_at
`

type UpdateCompetenceParams struct {
	CompetenceID string
	Code         string
	Name         string
	UpdatedAt    pgtype.Timestamp
}

func (q *Queries) UpdateCompetence(ctx context.Context, arg UpdateCompetenceParams) error {
	_, err := q.db.Exec(ctx, updateCompetence,
		arg.CompetenceID,
		arg.Code,
		arg.Name,
		arg.UpdatedAt,
	)
	return err
}
