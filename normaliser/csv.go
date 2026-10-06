package normaliser

import (
	"encoding/json"
	"encoding/csv"
	"os"
	"strings"
)

type CSVResult struct {
	RowsProcessed int
	Corrections   int
	ManualChecks  int
}

func LoadSpecies(filename string) ([]Species, error) {
	file, err := os.Open(filename)

	if err != nil {return nil, err}

	defer file.Close()

	var species []Species

	err = json.NewDecoder(file).Decode(&species)

	if err != nil {return nil, err}

	return species, nil
}

func ProcessCSV(inputFile string, outputFile string, columnName string, species []Species) (CSVResult, error) {
	input, err := os.Open(inputFile)

	if err != nil {return CSVResult{}, err}

	defer input.Close()

	reader := csv.NewReader(input)

	rows, err := reader.ReadAll()

	if err != nil {return CSVResult{}, err}

	if len(rows) == 0 {return CSVResult{}, nil}

	header := rows[0]

	columnIndex := -1

	for i, column := range header {
		if strings.EqualFold(column, columnName) {
			columnIndex = i
			break
		}
	}

	if columnIndex == -1 {return CSVResult{}, os.ErrNotExist}

	header = append(header, columnName+"_processed")

	output, err := os.Create(outputFile)

	if err != nil {return CSVResult{}, err}

	defer output.Close()

	writer := csv.NewWriter(output)
	defer writer.Flush()

	err = writer.Write(header)

	if err != nil {
		return CSVResult{}, err
	}

	result := CSVResult{}

	for _, row := range rows[1:] {
		if columnIndex >= len(row) {
			continue
		}

		inputName := row[columnIndex]

		normalisedName, found := NormaliseName(inputName, species)

		if found {
			row = append(row, normalisedName)

			if !strings.EqualFold(strings.TrimSpace(inputName), normalisedName) {
				result.Corrections++
			}
		} else {
			row = append(row, "ERROR: MANUAL CHECK")
			result.ManualChecks++
		}

		result.RowsProcessed++

		err = writer.Write(row)

		if err != nil {
			return result, err
		}
	}

	return result, nil
}