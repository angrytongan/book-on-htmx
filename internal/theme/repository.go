package theme

type Repository interface {
	Active(id int) (string, error)
	SetActive(id int, name string) error
	Themes(activeTheme string) ([]ThemeLink, error)
}
