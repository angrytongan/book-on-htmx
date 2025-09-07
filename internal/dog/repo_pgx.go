package dog

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
	sqlBreeds = `
		SELECT DISTINCT breed FROM dog ORDER BY breed
	`

	sqlColours = `
		SELECT DISTINCT colour FROM dog ORDER BY colour
	`

	sqlAll = `
		SELECT
			colour,
			breed,
			name
		FROM dog
			WHERE
				    ('all' = $1 OR colour = $1)
				AND ('all' = $2 OR breed = $2)
		ORDER BY %s
		LIMIT $3
	`
)

func NewPGXPoolRepository(pool *pgxpool.Pool) Repository {
	return &PGXPoolRepository{pool}
}

func basic[T any](p *PGXPoolRepository, ctx context.Context, q string) ([]T, error) {
	rows, _ := p.pool.Query(ctx, q)

	out, err := pgx.CollectRows(rows, pgx.RowTo[T])
	if err != nil {
		return out, fmt.Errorf("pgx.CollectRows(): %w", err)
	}

	return out, nil
}

func (p *PGXPoolRepository) Colours(ctx context.Context) ([]string, error) {
	out, err := basic[string](p, ctx, sqlColours)
	if err != nil {
		return out, fmt.Errorf("p.basic(sqlColours): %w", err)
	}

	return out, nil
}

func (p *PGXPoolRepository) Breeds(ctx context.Context) ([]string, error) {
	out, err := basic[string](p, ctx, sqlBreeds)
	if err != nil {
		return out, fmt.Errorf("p.basic(sqlBreeds): %w", err)
	}

	return out, nil
}

func (p *PGXPoolRepository) All(
	ctx context.Context,
	colour, breed string,
	orderBy string,
	limit int,
) ([]Dog, error) {

	q := fmt.Sprintf(sqlAll, orderBy)

	rows, _ := p.pool.Query(
		ctx,
		q,
		colour, // $1
		breed,  // $2
		limit,  // $4
	)

	out, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dog])
	if err != nil {
		return out, fmt.Errorf("pgx.CollectRows(): %w", err)
	}

	return out, nil
}
