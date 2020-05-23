package reviewers

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Doctype struct{}

func (doc Doctype) ReviewerName() string {
	return "Doctype Reviewer"
}

func (doc Doctype) Accepts(fileName string) bool {
	isPartial := strings.HasPrefix(fileName, "_")
	return !isPartial
}

func (doc Doctype) Review(page io.Reader) ([]Fault, error) {
	result := []Fault{}
	r := bufio.NewReader(page)

	var prevLine, line []byte
	var number int
	var err error

	for {
		number++

		line, _, err = r.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return result, err
		}

		if bytes.Contains(line, []byte("<html")) && !bytes.Contains(prevLine, []byte("<!DOCTYPE html>")) {
			result = append(result, Fault{
				ReviewerName: doc.ReviewerName(),
				LineNumber:   number,

				Rule: Rule{
					Name: "Missing Doctype",
					Code: "0001",
				},
			})
			break
		}

		prevLine = line
	}

	return result, nil
}
