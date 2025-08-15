package search

import (
	"sort"
	"strings"
)

func Term(term string) []string {
	terms := []string{
		"bulginess", "jocum", "censure", "Periploca", "unartificial",
		"overproportionately", "odontophorous", "presbyophrenia", "jaguar",
		"heteroplasm", "conchiferous", "artotype", "foothill", "rebesiege",
		"meadow", "forcibly", "rotundly", "crotyl", "classable",
		"conterminousness", "Guarnieri", "veratroyl", "Pandarus", "nonassenting",
		"rackett", "surfacy", "dioicous", "Salomonian", "quadrifolium",
		"multituberculated", "overtest", "leuco", "Salmo", "apheliotropically",
		"Chalcididae", "toxical", "imprecise", "impetre", "octopine", "neolater",
		"bonnily", "Chamaeleontidae", "Atrypa", "jumbuck", "unshameable",
		"provisive", "twitteringly", "pewdom", "synclinal", "butcherliness",
	}

	results := []string{}

	for _, t := range terms {
		if strings.Contains(t, term) {
			results = append(results, t)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i] < results[j] {
			return true
		}

		return false
	})

	return results
}
