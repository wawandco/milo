package reviewers

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type StyleAttribute struct{}

func (ol StyleAttribute) ReviewerName() string {
	return "attribute/style"
}

func (ol StyleAttribute) Accepts(filePath string) bool {
	return true
}

func (ol StyleAttribute) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return result, err
	}

	doc.Find("[style]").Each(func(i int, s *goquery.Selection) {
		for _, node := range s.Nodes {
			result = append(result, Fault{
				Reviewer: ol.ReviewerName(),
				Line:     node.Line,
				Path:     path,
				Rule:     Rules["0009"],
			})
		}
	})

	return result, nil
}
