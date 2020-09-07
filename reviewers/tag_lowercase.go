package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

// StyleTag is a reviewer that checks that all tags into the HTML file must be in lowercase.
type TagLowercase struct{}

// ReviewerName returns the reviewer name.
func (t TagLowercase) ReviewerName() string {
	return "tag/lowercase"
}

// Accepts checks if the file can be reviewed.
func (t TagLowercase) Accepts(path string) bool {
	return true
}

// Review returns a fault for each tag that is not in lowercase.
// The faults will be added if the HTML file have tags like <DIV>, <Input>, etc.
func (t TagLowercase) Review(path string, page io.Reader) ([]Fault, error) {
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

		if tt != html.EndTagToken && tt != html.StartTagToken && tt != html.SelfClosingTagToken {
			continue
		}

		tok := z.Token()
		if tok.Data == tok.Name {
			continue
		}

		result = append(result, Fault{
			Reviewer: t.ReviewerName(),
			Line:     tok.Line,
			Col:      tok.Col,
			Path:     path,
			Rule:     Rules[t.ReviewerName()],
		})
	}

	return result, nil
}
