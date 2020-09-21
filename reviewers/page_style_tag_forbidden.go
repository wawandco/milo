package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
	"github.com/wawandco/milo/external/html/atom"
)

// PageStyleTagForbidden is a reviewer that checks if the HTML file has style tags.
type PageStyleTagForbidden struct{}

// ReviewerName returns the reviewer name.
func (css PageStyleTagForbidden) ReviewerName() string {
	return "page-style-tag-forbidden"
}

// Accepts checks if the file can be reviewed.
func (css PageStyleTagForbidden) Accepts(path string) bool {
	return true
}

// Review returns a fault for each <style> tag found in HTML file.
func (css PageStyleTagForbidden) Review(path string, page io.Reader) ([]Fault, error) {
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

		if tt != html.StartTagToken && tt != html.SelfClosingTagToken {
			continue
		}

		tok := z.Token()
		if tok.DataAtom != atom.Style {
			continue
		}

		result = append(result, Fault{
			Reviewer: css.ReviewerName(),
			Line:     tok.Line,
			Col:      tok.Col,
			Path:     path,
			Rule:     Rules[css.ReviewerName()],
		})
	}

	return result, nil
}
