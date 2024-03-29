package reviewers

import (
	"io"
	"strings"

	"github.com/wawandco/milo/internal/html"
)

// AttributeAltRequired is a reviewer that checks that all img tags have alt attribute.
type AttributeAltRequired struct{}

// ReviewerName returns the reviewer name.
func (at AttributeAltRequired) ReviewerName() string {
	return "attribute-alt-required"
}

// Accepts checks if the file can be reviewed.
func (at AttributeAltRequired) Accepts(filePath string) bool {
	return true
}

// Review returns a fault for each img tag that does not have the alt attribute.
func (at AttributeAltRequired) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		token := z.Token()
		if !at.tagRequiresAlt(token) || at.hasAlt(token) {
			continue
		}

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			result = append(result, Fault{
				Reviewer: at.ReviewerName(),
				Line:     token.Line,
				Col:      token.Col,
				Path:     path,
				Rule:     Rules[at.ReviewerName()],
			})
		}
	}

	return result, nil
}

func (at AttributeAltRequired) tagRequiresAlt(token html.Token) bool {
	switch token.DataAtom.String() {
	case "img":
		return true
	case "input":
		for _, attr := range token.Attr {
			if attr.Key != "type" || strings.ToLower(attr.Val) != "image" {
				continue
			}

			return true
		}

		return false
	case "area":
		for _, attr := range token.Attr {
			if attr.Key == "href" {
				return true
			}
		}

		return true
	default:
		return false
	}
}

func (at AttributeAltRequired) hasAlt(token html.Token) bool {
	for _, attr := range token.Attr {
		if attr.Key == "alt" && attr.Val != "" {
			return true
		}
	}

	return false
}
