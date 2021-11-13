package csvwriter

import (
	"encoding/csv"
	"os"
)

type RowHeaderGetter interface {
	CSVHeader() []string
}

type RowBodyGetter interface {
	CSVRow() []string
}

type MultiRowBodyGetter interface {
	CSVRows() [][]string
}

type FullFileWriter interface {
	RowHeaderGetter
	MultiRowBodyGetter
}

type CSVWriter struct {
	w        *csv.Writer
	filename string
}

func (c *CSVWriter) WriteHeader(g RowHeaderGetter) error {
	err := c.w.Write(g.CSVHeader())
	if err != nil {
		err = wrapError(err, "failed writing a header row to file: %s", c.filename)
	}
	return err
}

func (c *CSVWriter) WriteRow(g RowBodyGetter) error {
	err := c.w.Write(g.CSVRow())
	if err != nil {
		err = wrapError(err, "failed writing a header row to file: %s", c.filename)
	}
	return err
}

func (c *CSVWriter) WriteFull(g FullFileWriter) error {
	err := c.WriteHeader(g)
	if err != nil {
		return err
	}
	err = c.w.WriteAll(g.CSVRows())
	if err != nil {
		err = wrapError(err, "failed writing multiple rows to file: %s", c.filename)
	}
	return err
}

// Creates new file to write to. Return error if file already exist.
func NewCSVWriterToNewFile(filename string) (w *CSVWriter, err error) {
	file, err := os.Create(filename)
	if err != nil {
		err = wrapError(err, "failed to create new file: %s", filename)
		return
	}

	return &CSVWriter{csv.NewWriter(file), filename}, nil
}
