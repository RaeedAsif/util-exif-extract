package service

import (
	"encoding/csv"
	"fmt"
	"os"
)

var cols = []string{"path", "latitude", "longitude"}
var errCols = []string{"path", "error"}

type Writer struct {
	file      *os.File
	errFile   *os.File
	writer    *csv.Writer
	errWriter *csv.Writer
}

// NewWriter creates a new Writer
func NewWriter(filename string) (*Writer, error) {
	outDir := "output/"
	path := outDir + filename
	errPath := outDir + "error_" + filename

	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err := os.Mkdir(outDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(file)

	err = writer.Write(cols)
	if err != nil {
		return nil, err
	}

	errfile, err := os.Create(errPath)
	if err != nil {
		return nil, err
	}

	errwriter := csv.NewWriter(errfile)

	err = errwriter.Write(errCols)
	if err != nil {
		return nil, err
	}

	return &Writer{
		file:      file,
		errFile:   errfile,
		writer:    writer,
		errWriter: errwriter,
	}, nil
}

// Write a row of data
func (e *Writer) WriteRow(path string, latitude float64, longitude float64) error {
	err := e.writer.Write([]string{path, fmt.Sprint(latitude), fmt.Sprint(longitude)})
	if err != nil {
		return err
	}

	return nil
}

// WriteErrorRow writes a row of data with an error
func (e *Writer) WriteErrorRow(path string, err error) error {
	err = e.errWriter.Write([]string{path, err.Error()})
	if err != nil {
		return err
	}

	return nil
}

// Flush the writer and close the file
func (e *Writer) Close() {
	e.writer.Flush()
	e.errWriter.Flush()
	e.file.Close()
	e.errFile.Close()
}
