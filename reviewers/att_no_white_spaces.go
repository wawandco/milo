package reviewers

import (
	"io"
	"regexp"

	"github.com/wawandco/milo/external/html"
)

// AttrNoWhiteSpaces is a reviewer that checks that there is not a blank spaces between tag attribute and its value.
type AttrNoWhiteSpaces struct{}

// ReviewerName returns the reviewer name.
func (at AttrNoWhiteSpaces) ReviewerName() string {
	return "attribute/no-white-spaces"
}

// Accepts checks if the file can be reviewed.
func (at AttrNoWhiteSpaces) Accepts(path string) bool {
	return true
}

// Review returns a fault for each tag that has blank spaces between attribute and value
// For example.
//  <div class="foo"> is valid.
//  <div class= "foo"> will return a fault.
func (at AttrNoWhiteSpaces) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}
	exp := regexp.MustCompile(`\S+(\s+=\s*|\s*=\s+)`)

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()

		if err := z.Err(); err != nil {
			if err == io.EOF {
				break
			}

			return []Fault{}, err
		}

		if tt != html.StartTagToken && tt != html.SelfClosingTagToken {
			continue
		}

		tok := z.Token()
		if exp.MatchString(string(z.Raw())) {
			result = append(result, Fault{
				Reviewer: at.ReviewerName(),
				Line:     tok.Line,
				Col:      tok.Col,
				Path:     path,
				Rule:     Rules[at.ReviewerName()],
			})
		}
	}

	return result, nil
}
