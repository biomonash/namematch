package normaliser

import "strings"

type Species struct {
	Name         string   `json:"name"`
	Alternatives []string `json:"alternatives"`
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

func EditDistance(a string, b string) int {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	distance := make([][]int, len(a)+1)

	for i := range distance {
		distance[i] = make([]int, len(b)+1)
	}

	for i := 0; i <= len(a); i++ {
		distance[i][0] = i
	}

	for j := 0; j <= len(b); j++ {
		distance[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			cost := 0

			if a[i-1] != b[j-1] {
				cost = 1
			}

			insert := distance[i][j-1] + 1
			delete := distance[i-1][j] + 1
			replace := distance[i-1][j-1] + cost

			distance[i][j] = min(insert, delete, replace)
		}
	}

	return distance[len(a)][len(b)]
}

func min(a int, b int, c int) int {
	if a < b && a < c {
		return a
	}

	if b < c {
		return b
	}

	return c
}

// NormaliseName tries to find the standard species name. If the name is not an exact match, it checks for small typos.
func NormaliseName(input string, species []Species) (string, bool) {
	input = strings.TrimSpace(input)

	// Check exact names and alternatives first.
	result, found := FindSpecies(input, species)

	if found {
		return result, true
	}

	// Check for small spelling mistakes.
	best_match := ""
	best_distance := 3

	for _, s := range species {
		distance := EditDistance(input, s.Name)

		if distance < best_distance {
			best_distance = distance
			best_match = s.Name
		}

		for _, alternative := range s.Alternatives {
			distance := EditDistance(input, alternative)

			if distance < best_distance {
				best_distance = distance
				best_match = s.Name
			}
		}
	}

	if best_match != "" {
		return best_match, true
	}

	return "", false
}