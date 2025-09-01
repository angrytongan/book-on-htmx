package nav

import "slices"

type PageLink struct {
	Href   string
	Label  string
	Active bool
}

var (
	pageLinks = []PageLink{
		{"/", "Home", false},
		{"/dashboard", "Dashboard", false},
		{"/leaflet", "Leaflet", false},
		{"/search", "Search", false},
		{"/settings", "Settings", false},
		{"/tabs", "Tabs", false},
		{"/toast", "Toast", false},
		{"/theme", "Theme", false},
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
