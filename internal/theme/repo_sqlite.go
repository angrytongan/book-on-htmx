package theme

import (
	"database/sql"
	"fmt"
)

const (
	sqlThemeByID = `
		SELECT theme FROM settings WHERE id = $1;
	`

	sqlSetThemeByID = `
		INSERT INTO settings (id, theme) VALUES ($1, $2)
		ON CONFLICT(id) DO UPDATE SET theme = excluded.theme;
	`

	sqlThemes = `
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

type ThemeRepository struct {
	db *sql.DB
}

func SQLRepository(db *sql.DB) Repository {
	tr := &ThemeRepository{db}

	return tr
}

func (tr *ThemeRepository) Active(id int) (string, error) {
	var theme string

	row := tr.db.QueryRow(sqlThemeByID, id)
	if err := row.Scan(&theme); err != nil {
		return "light", fmt.Errorf("row.Scan(): %w", err) // default to "light"
	}

	return theme, nil
}

func (tr *ThemeRepository) SetActive(id int, name string) error {
	_, err := tr.db.Exec(sqlSetThemeByID, id, name)
	if err != nil {
		return fmt.Errorf("tr.db.Exec(sqlSetThemeByID, %d, %s): %w", id, name, err)
	}

	return nil
}

func (tr *ThemeRepository) Themes(activeTheme string) ([]ThemeLink, error) {
	rows, err := tr.db.Query(sqlThemes, activeTheme)
	if err != nil {
		return []ThemeLink{}, fmt.Errorf("tr.db.Query(sqlThemes, %s): %w", activeTheme, err)
	}

	links := []ThemeLink{}

	for rows.Next() {
		var l ThemeLink

		if err := rows.Scan(&l.Label, &l.Value, &l.Active); err != nil {
			return links, fmt.Errorf("rows.Scan(): %w", err)
		}

		links = append(links, l)
	}

	return links, nil
}
