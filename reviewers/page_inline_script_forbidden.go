package reviewers

import (
	"io"
	"regexp"

	"github.com/wawandco/milo/internal/html"
)

// PageInlineScriptForbidden is a reviewer that checks that tags does not have inline javascript actions.
type PageInlineScriptForbidden struct{}

// ReviewerName returns the reviewer name.
func (i PageInlineScriptForbidden) ReviewerName() string {
	return "page-inline-script-forbidden"
}

// Accepts checks if the file can be reviewed.
func (i PageInlineScriptForbidden) Accepts(path string) bool {
	return true
}

// Review return a fault for each tag that has an inline event attribute.
// For example, <button ... onclick="foo();">.
func (i PageInlineScriptForbidden) Review(path string, page io.Reader) ([]Fault, error) {
	var fault []Fault
	var onEventRegexp = regexp.MustCompile(`(?i)^on(unload|message|submit|select|scroll|resize|mouseover|mouseout|mousemove|mouseleave|mouseenter|mousedown|load|keyup|keypress|keydown|focus|dblclick|click|change|blur|error)$`)
	var javascriptRegexp = regexp.MustCompile(`(?i)^\s*javascript:`)

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
			if onEventRegexp.MatchString(attr.Key) {
				fault = append(fault, Fault{
					Reviewer: i.ReviewerName(),
					Line:     tok.Line,
					Col:      attr.Col,
					Path:     path,
					Rule:     Rules[i.ReviewerName()],
				})

				continue
			}

			if (attr.Key == "src" || attr.Key == "href") &&
				javascriptRegexp.MatchString(attr.Val) {
				fault = append(fault, Fault{
					Reviewer: i.ReviewerName(),
					Line:     tok.Line,
					Col:      attr.Col,
					Path:     path,
					Rule:     Rules[i.ReviewerName()],
				})
			}
		}
	}

	return fault, nil
}
