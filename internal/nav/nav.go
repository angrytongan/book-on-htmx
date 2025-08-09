package nav

type PageLink struct {
	Href  string
	Label string
}

func PageLinks() []PageLink {
	return []PageLink{
		{
			Href:  "/",
			Label: "Home",
		},
		{
			Href:  "/tabs",
			Label: "Tabs",
		},
		{
			Href:  "/settings",
			Label: "Settings",
		},
	}
}
