package main

import (
	"fmt"
	"species_name_normaliser/normaliser"
)

func main() {
	species := []normaliser.Species{
		{Name: "Panthera tigris", Alternatives: []string{"tiger"}},
		{Name: "Canis lupus", Alternatives: []string{"wolf"}},
		{Name: "Balaenoptera musculus", Alternatives: []string{"blue whale"}},
		{Name: "Orcinus orca", Alternatives: []string{"orca", "killer whale"}},
	}

	names := []string{
		"Panthera tigris",
		"PANTHERA tigris",
		"tiger",
		"wolf",
		"tIgEr",
		"blue wale",
		"orca",
		"killer whale",
	}

	for _, name := range names {
		result, found := normaliser.FindSpecies(name, species)

		if found {
			fmt.Println(name, "->", result)
		} else {
			fmt.Println(name, "-> no match")
		}
	}
}

