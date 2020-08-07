package reviewers

import (
	"io"
	"strings"

	"github.com/wawandco/milo/external/html"
)

type AltRequired struct{}

func (at AltRequired) ReviewerName() string {
	return "attribute/alt-required"
}

func (at AltRequired) Accepts(filePath string) bool {
	return true
}

func (at AltRequired) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			token := z.Token()

			if !at.tagRequiresAlt(token) || at.hasAlt(token) {
				continue
			}

			result = append(result, Fault{
				Reviewer: at.ReviewerName(),
				Line:     token.Line,
				Path:     path,
				Rule:     Rules[at.ReviewerName()],
			})
		}
	}

	return result, nil
}

func (at AltRequired) tagRequiresAlt(token html.Token) bool {
	if token.DataAtom.String() == "img" {
		return true
	}

	// input[type=image]
	if token.DataAtom.String() == "input" {
		for _, attr := range token.Attr {
			if attr.Key != "type" || strings.ToLower(attr.Val) != "image" {
				continue
			}

			return true
		}

		return false
	}

	// area[href]
	if token.DataAtom.String() == "area" {
		for _, attr := range token.Attr {
			if attr.Key == "href" {
				return true
			}
		}

		return true
	}

	return false
}

func (at AltRequired) hasAlt(token html.Token) bool {
	for _, attr := range token.Attr {
		if attr.Key == "alt" && attr.Val != "" {
			return true
		}
	}

	return false
}
