package reviewers

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/wawandco/milo/external/html"
)

// DoctypePresent is a reviewer that checks if an HTML file has the <!DOCTYPE x> tag.
type DoctypePresent struct{}

// ReviewerName returns the reviewer name.
func (doc DoctypePresent) ReviewerName() string {
	return "doctype/present"
}

// Accepts checks if the file can be reviewed.
func (doc DoctypePresent) Accepts(filePath string) bool {
	fileName := filepath.Base(filePath)
	isPartial := strings.HasPrefix(fileName, "_")

	return !isPartial
}

// Review returns a fault if HTML file does not have the <!DOCTYPE x> tag.
// For html files that do not have the <html> tag, the review will not return any fault.
func (doc DoctypePresent) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}
	var (
		found   bool
		htmlTag *html.Token
	)

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		token := z.Token()
		if token.DataAtom.String() == "html" {
			htmlTag = &token

			continue
		}

		if tt == html.DoctypeToken {
			found = true
		}
	}

	if htmlTag == nil || found {
		return result, nil
	}

	result = append(result, Fault{
		Reviewer: doc.ReviewerName(),
		Line:     htmlTag.Line,
		Col:      htmlTag.Col,
		Path:     path,
		Rule:     Rules[doc.ReviewerName()],
	})

	return result, nil
}
