package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

// PageInlineCSSForbidden is a reviewer that checks that all tags in HTML file do not have the style attribute.
type PageInlineCSSForbidden struct{}

// ReviewerName returns the reviewer name.
func (css PageInlineCSSForbidden) ReviewerName() string {
	return "page-inline-css-forbidden"
}

// Accepts checks if the file can be reviewed.
func (css PageInlineCSSForbidden) Accepts(path string) bool {
	return true
}

// Review returns a list of faults for each tag that has the style attribute.
// It returns faults if there are tags like <div ... style="color: red;">.
func (css PageInlineCSSForbidden) Review(path string, page io.Reader) ([]Fault, error) {
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
		for _, attr := range tok.Attr {
			if attr.Key == "style" {
				result = append(result, Fault{
					Reviewer: css.ReviewerName(),
					Line:     attr.Line,
					Col:      attr.Col,
					Path:     path,
					Rule:     Rules[css.ReviewerName()],
				})
			}
		}
	}

	return result, nil
}
