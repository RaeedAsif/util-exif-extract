package service

import (
	"encoding/csv"
	"os"
)

// GetRecords returns a new Reader instance
func GetRecords(filename string) ([][]string, error) {
	outDir := "output/"
	path := outDir + filename

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	return reader.ReadAll()
}
