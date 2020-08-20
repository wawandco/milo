package reviewers

import (
	"fmt"
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

		token := z.Token()
		if !at.tagRequiresAlt(token) || at.hasAlt(token) {
			continue
		}

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			fmt.Println("Adding Fault")

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
		fmt.Println("IMG")
		return true
	}

	// input[type=image]
	if token.DataAtom.String() == "input" {
		for _, attr := range token.Attr {
			if attr.Key != "type" || strings.ToLower(attr.Val) != "image" {
				continue
			}

			fmt.Println("Input image")

			return true
		}

		return false
	}

	// area[href]
	if token.DataAtom.String() == "area" {
		for _, attr := range token.Attr {
			if attr.Key == "href" {
				fmt.Println("Area")

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
