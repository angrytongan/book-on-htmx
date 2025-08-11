package theme

import "context"

type Repository interface {
	Active(ctx context.Context, id int) (string, error)
	SetActive(ctx context.Context, id int, name string) error
	Themes(ctx context.Context, activeTheme string) ([]ThemeLink, error)
}
