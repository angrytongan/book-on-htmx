package search

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGXPoolRepository struct {
	pool *pgxpool.Pool
}

const (
	sqlTerm = `
		SELECT
			word
		FROM
			word
		WHERE
			word ILIKE '%' || $1 || '%'
	`
)

func NewPGXPoolRepository(pool *pgxpool.Pool) Repository {
	return &PGXPoolRepository{pool}
}

func (sr *PGXPoolRepository) Term(ctx context.Context, term string) ([]string, error) {
	rows, _ := sr.pool.Query(ctx, sqlTerm, term)
	words, err := pgx.CollectRows(rows, pgx.RowTo[string])

	if err != nil {
		return words, fmt.Errorf("pgx.CollectRows(): %w", err)
	}

	return words, nil
}
