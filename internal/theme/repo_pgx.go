package theme

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	pgxThemeByID = `
		SELECT theme FROM settings WHERE id = $1;
	`

	pgxSetThemeByID = `
		UPDATE settings SET theme = $2 WHERE id = $1;
	`

	pgxThemes = `
		SELECT
			label,
			value,
			value = $1 AS active
		FROM
			themes
		ORDER BY
			label
	`
)

type PGXPoolRepository struct {
	pool *pgxpool.Pool
}

func NewPGXPoolRepository(pool *pgxpool.Pool) Repository {
	return &PGXPoolRepository{pool}
}

func (tr *PGXPoolRepository) Active(ctx context.Context, id int) (string, error) {
	var activeTheme string

	row := tr.pool.QueryRow(ctx, pgxThemeByID, id)
	if err := row.Scan(&activeTheme); err != nil {
		return "", fmt.Errorf("row.Scan(): %w", err)
	}

	return activeTheme, nil
}

func (tr *PGXPoolRepository) SetActive(ctx context.Context, id int, name string) error {
	_, err := tr.pool.Exec(ctx, pgxSetThemeByID, id, name)
	if err != nil {
		return fmt.Errorf("tr.pool.Exec(pgxSetThemeByID, %d, %s): %w", id, name, err)
	}

	return nil
}

func (tr *PGXPoolRepository) Themes(ctx context.Context, activeTheme string) ([]ThemeLink, error) {
	rows, _ := tr.pool.Query(ctx, pgxThemes, activeTheme)
	themes, err := pgx.CollectRows(rows, pgx.RowToStructByName[ThemeLink])
	if err != nil {
		return []ThemeLink{}, fmt.Errorf("pgx.CollectRows(): %w", err)
	}

	return themes, nil
}
