package reviewers

import (
	"io"
	"io/ioutil"
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

	bcontent, err := ioutil.ReadAll(page)
	if err != nil {
		return result, err
	}

	content := string(bcontent)
	if !strings.Contains(strings.ToLower(content), "<html") {
		return result, nil
	}

	lines := strings.Split(content, "\n")
	for number, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		for _, valid := range validDoctypes {
			if strings.Contains(line, strings.ToLower(valid)) {
				return result, nil
			}
		}

		result = append(result, Fault{
			Reviewer: doc.ReviewerName(),
			Line:     number + 1,
			Path:     path,

			Rule: Rules["0002"],
		})

		break
	}

	return result, nil
}
