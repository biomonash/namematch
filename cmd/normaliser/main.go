package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"species_name_normaliser/normaliser"
)

func main() {
	input_file := flag.String("input", "", "CSV file to process")
	output_file := flag.String("output", "processed.csv", "output CSV file")
	column := flag.String("column", "", "CSV column containing species names")

	flag.Parse()

	species_file, err := os.Open("config/species.json")

	if err != nil {
		fmt.Println("Could not open species config:", err)
		return
	}

	defer species_file.Close()

	var species []normaliser.Species

	err = json.NewDecoder(species_file).Decode(&species)

	if err != nil {
		fmt.Println("Could not read species config:", err)
		return
	}

	if *input_file == "" || *column == "" {
		fmt.Println("Please provide an input CSV and column name.")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("go run ./cmd/normaliser -input temp_test.csv -column species")
		return
	}

	result, err := normaliser.ProcessCSV(
		*input_file,
		*output_file,
		*column,
		species,
	)

	if err != nil {
		fmt.Println("Could not process CSV:", err)
		return
	}

	fmt.Println("CSV processing complete.")
	fmt.Println("Rows processed:", result.RowsProcessed)
	fmt.Println("Corrections:", result.Corrections)
	fmt.Println("Manual checks:", result.ManualChecks)
	fmt.Println("Output:", *output_file)
}

// go run ./cmd/normaliser -input temp_test.csv -column species