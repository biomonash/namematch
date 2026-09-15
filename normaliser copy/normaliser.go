package normaliser

import "strings"

type Species struct {
	Name         string
	Alternatives []string
}

func FindSpecies(input string, species []Species) (string, bool) {
	for _, s := range species {
		if strings.EqualFold(input, s.Name) {
			return s.Name, true
		}

		for _, alternative := range s.Alternatives {
			if strings.EqualFold(input, alternative) {
				return s.Name, true
			}
		}
	}

	return "", false
}