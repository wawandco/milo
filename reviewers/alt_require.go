package reviewers

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AltRequired struct{}

func (at AltRequired) ReviewerName() string {
	return "ol-ul/valid"
}

func (at AltRequired) Accepts(filePath string) bool {
	return true
}

func (at AltRequired) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return result, err
	}

	matched := doc.Find("area[href], input[type=image], img")

	for _, node := range matched.Nodes {
		found := false
		for _, attr := range node.Attr {
			if strings.ToLower(attr.Key) == "alt" && attr.Key != "" {
				found = true
				break
			}
		}

		if found {
			continue
		}

		result = append(result, Fault{
			Reviewer: at.ReviewerName(),
			Line:     node.Line,
			Path:     path,
			Rule:     Rules["0012"],
		})
	}

	return result, nil
}
