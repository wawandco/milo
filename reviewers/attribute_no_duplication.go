package reviewers

import (
	"io"

	"github.com/wawandco/milo/internal/html"
)

// AttributeNoDuplication is a reviewer that checks that tags does not have an attribute duplicated.
type AttributeNoDuplication struct{}

// ReviewerName returns the reviewer name.
func (a AttributeNoDuplication) ReviewerName() string {
	return "attribute-no-duplication"
}

// Accepts checks if the file can be reviewed.
func (a AttributeNoDuplication) Accepts(path string) bool {
	return true
}

// Review returns a fault for each thag that has duplicated attributes
// For expample, <div class="..." data-attr=".." class="..."> a fault will be added because this tag has class attribute duplicated.
func (a AttributeNoDuplication) Review(path string, page io.Reader) ([]Fault, error) {
	var fault []Fault
	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		tok := z.Token()
		keys := map[string]bool{}
		for _, attr := range tok.Attr {
			if keys[attr.Key] {
				fault = append(fault, Fault{
					Reviewer: a.ReviewerName(),
					Line:     tok.Line,
					Col:      attr.Col,
					Path:     path,
					Rule:     Rules[a.ReviewerName()],
				})
			}

			keys[attr.Key] = true
		}
	}

	return fault, nil
}
