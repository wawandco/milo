package reviewers

import (
	"io"
	"strings"

	"github.com/wawandco/milo/external/html"
)

var (
	// booleanAttributes list of attributes that not need a value.
	booleanAttributes = []string{
		"allowfullscreen",
		"allowpaymentrequest",
		"async",
		"autofocus",
		"autoplay",
		"checked",
		"contenteditable",
		"controls",
		"default",
		"defer",
		"disabled",
		"formnovalidate",
		"frameborder",
		"hidden",
		"ismap",
		"itemscope",
		"loop",
		"multiple",
		"muted",
		"nomodule",
		"novalidate",
		"open",
		"playsinline",
		"readonly",
		"required",
		"reversed",
		"scoped",
		"selected",
		"typemustmatch",
	}
)

// AttrValueNotEmpty is a reviewer that checks that all tags have a value.
type AttrValueNotEmpty struct{}

// ReviewerName returns the reviewer name.
func (a AttrValueNotEmpty) ReviewerName() string {
	return "attribute/value-not-empty"
}

// Accepts checks if the file can be reviewed.
func (a AttrValueNotEmpty) Accepts(path string) bool {
	return true
}

// Review return a fault for each a tag has an empty attribute value like: <div class=""...>.
func (a AttrValueNotEmpty) Review(path string, page io.Reader) ([]Fault, error) {
	var fault []Fault
	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		tok := z.Token()
		for _, attr := range tok.Attr {
			if _contains(attr.Name) {
				continue
			}

			if strings.TrimSpace(attr.Val) == "" {
				fault = append(fault, Fault{
					Reviewer: a.ReviewerName(),
					Line:     tok.Line,
					Col:      attr.Col,
					Path:     path,
					Rule:     Rules[a.ReviewerName()],
				})
			}
		}
	}

	return fault, nil
}

func _contains(attr string) bool {
	for _, a := range booleanAttributes {
		if strings.EqualFold(a, attr) {
			return true
		}
	}

	return false
}
