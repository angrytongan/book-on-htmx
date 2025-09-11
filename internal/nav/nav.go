package nav

import "slices"

type PageLink struct {
	Href   string
	Label  string
	Icon   string
	Active bool
}

var (
	pageLinks = []PageLink{
		{"/", "Home", "home", false},
		{"/dashboard", "Dashboard", "dashboard", false},
		{"/dog", "Dog", "pets", false},
		{"/drawer", "Drawer", "menu", false},
		{"/leaflet", "Leaflet", "map", false},
		{"/repl", "REPL", "terminal", false},
		{"/search", "Search", "search", false},
		{"/settings", "Settings", "settings", false},
		{"/tabs", "Tabs", "tab", false},
		{"/toast", "Toast", "notifications", false},
		{"/theme", "Theme", "palette", false},
	}
)

func PageLinks(activeHref string) []PageLink {
	pl := slices.Clone(pageLinks)

	for i, p := range pl {
		if p.Href == activeHref {
			pl[i].Active = true

			break
		}
	}

	return pl
}

func IsNavLink(href string) bool {
	for _, l := range pageLinks {
		if l.Href == href {
			return true
		}
	}

	return false
}
