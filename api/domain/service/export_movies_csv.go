package service

import (
	"bytes"
	"encoding/csv"
	"math/rand"
	"strconv"
)

func ExportMoviesCsv() (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}
	writer := csv.NewWriter(buffer)

	defer writer.Flush()
	err := writer.Write([]string{"name", "genre", "assessment", "synapses"})
	for i := 0; i < 1000000; i++ {

		// err check
		err = writer.Write([]string{"BAD BOYS", "comedy", strconv.Itoa(rand.Intn(100)), "fine"})
		if err != nil {
			return nil, err
		}
	}

	return buffer, err
}
