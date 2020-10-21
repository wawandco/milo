package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

// AttributeLowercase is a reviewer that checks that all tags attributes are in lowercase.
type AttributeLowercase struct{}

// ReviewerName returns the reviewer name.
func (a AttributeLowercase) ReviewerName() string {
	return "attribute-lowercase"
}

// Accepts checks if the file can be reviewed.
func (a AttributeLowercase) Accepts(path string) bool {
	return true
}

//  Review returns a fault for each tag attribute that is in lowercase.
// For example, <div CLASS="..."> or <div Class="..."> will return a fault.
func (a AttributeLowercase) Review(path string, page io.Reader) ([]Fault, error) {
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

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			tok := z.Token()
			for _, attr := range tok.Attr {
				if attr.Name != attr.Key {
					result = append(result, Fault{
						Reviewer: a.ReviewerName(),
						Line:     tok.Line,
						Col:      attr.Col,
						Rule:     Rules[a.ReviewerName()],
						Path:     path,
					})
				}
			}
		}
	}

	return result, nil
}
