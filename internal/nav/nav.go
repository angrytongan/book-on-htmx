package nav

type PageLink struct {
	Href   string
	Label  string
	Active bool
}

func PageLinks(activeHref string) []PageLink {
	pageLinks := []PageLink{
		{"/", "Home", false},
		{"/tabs", "Tabs", false},
		{"/settings", "Settings", false},
	}

	for i, p := range pageLinks {
		if p.Href == activeHref {
			pageLinks[i].Active = true

			break
		}
	}

	return pageLinks
}
