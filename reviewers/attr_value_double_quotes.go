package reviewers

import (
	"io"
	"strings"

	"github.com/wawandco/milo/external/html"
)

// AttrValueDoubleQuotes is a reviewer that checks that tag values are enclosed in double quotes.
type AttrValueDoubleQuotes struct{}

// ReviewerName returns the reviewer name.
func (a AttrValueDoubleQuotes) Accepts(path string) bool {
	return true
}

// Accepts checks if the file can be reviewed.
func (a AttrValueDoubleQuotes) ReviewerName() string {
	return "attribute/double-quotes"
}

// Review returns a fault for each attribute value that is not enclosed by double quotes.
// For example:
//
// <div class="foo" ...> is correct.
// <div class=foo ...>  will return a fault.
func (a AttrValueDoubleQuotes) Review(path string, page io.Reader) ([]Fault, error) {
	var fault []Fault

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()

		if err := z.Err(); err != nil {
			if err == io.EOF {
				break
			}

			return []Fault{}, nil
		}

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			token := z.Token()
			for _, attr := range token.Attr {
				if attr.Quote == "" && strings.TrimSpace(attr.Val) == "" {
					continue
				}

				if attr.Quote != "\"" {
					fault = append(fault, Fault{
						Reviewer: a.ReviewerName(),
						Line:     token.Line,
						Col:      attr.Col,
						Path:     path,
						Rule:     Rules[a.ReviewerName()],
					})
				}
			}
		}
	}

	return fault, nil
}
