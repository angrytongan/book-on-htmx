package nav

type PageLink struct {
	Href   string
	Label  string
	Active bool
}

func PageLinks(activeHref string) []PageLink {
	pageLinks := []PageLink{
		{"/", "Home", false},
		{"/leaflet", "Leaflet", false},
		{"/search", "Search", false},
		{"/settings", "Settings", false},
		{"/tabs", "Tabs", false},
		{"/toast", "Toast", false},
		{"/theme", "Theme", false},
	}

	for i, p := range pageLinks {
		if p.Href == activeHref {
			pageLinks[i].Active = true

			break
		}
	}

	return pageLinks
}
