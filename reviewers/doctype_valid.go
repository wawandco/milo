package reviewers

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/wawandco/milo/external/html"
)

var validDoctypes = []string{
	`<!DOCTYPE html>`,
	`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">`,
	`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`,
}

// DoctypePresent is a reviewer that checks if the HTML file has a valid DOCTYPE tag.
type DoctypeValid struct{}

// ReviewerName returns the reviewer name.
func (doc DoctypeValid) ReviewerName() string {
	return "doctype/valid"
}

// Accepts checks if the file can be reviewed.
func (doc DoctypeValid) Accepts(filePath string) bool {
	fileName := filepath.Base(filePath)
	isPartial := strings.HasPrefix(fileName, "_")

	return !isPartial
}

// Review returns a fault if HTML file does not have a valid DOCTYPE tag.
//
// The valids DOCTYPE are:
// 	- <!DOCTYPE html>
//  - <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
// 	- <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
//
// For html files that do not have the <html> tag, the review will not return any fault.
func (doc DoctypeValid) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	z := html.NewTokenizer(page)

	for {
		tt := z.Next()
		if err := z.Err(); err != nil {
			if err == io.EOF {
				break
			}

			return []Fault{}, err
		}

		if tt != html.DoctypeToken {
			continue
		}

		tok := z.Token()
		for _, valid := range validDoctypes {
			if strings.Contains(strings.ToLower(string(z.Raw())), strings.ToLower(valid)) {
				return result, nil
			}
		}

		result = append(result, Fault{
			Reviewer: doc.ReviewerName(),
			Line:     tok.Line,
			Col:      tok.Col,
			Path:     path,

			Rule: Rules[doc.ReviewerName()],
		})
	}

	return result, nil
}
