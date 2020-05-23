package reviewers

import (
	"bufio"
	"io"
	"path/filepath"
	"strings"
)

var validDoctypes = []string{
	"<!DOCTYPE html>",
	`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">`,
	`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`,
}

type DoctypeValid struct{}

func (doc DoctypeValid) ReviewerName() string {
	return "doctype/valid"
}

func (doc DoctypeValid) Accepts(filePath string) bool {
	fileName := filepath.Base(filePath)
	isPartial := strings.HasPrefix(fileName, "_")
	return !isPartial
}

func (doc DoctypeValid) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}
	var number int
	var line string

	scanner := bufio.NewScanner(page)
	for scanner.Scan() {
		number++

		line = scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		lineLower := strings.ToLower(line)

		for _, valid := range validDoctypes {
			if strings.Contains(lineLower, strings.ToLower(valid)) {
				return result, nil
			}
		}

		result = append(result, Fault{
			Reviewer: doc.ReviewerName(),
			Line:     number,
			Path:     path,

			Rule: Rules["0002"],
		})

		break
	}

	return result, nil
}
