package theme

type ThemeLink struct {
	Label  string
	Value  string
	Active bool
}

func DefaultTheme() string {
	return ""
}
