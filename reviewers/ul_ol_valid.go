package reviewers

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type OlValid struct{}

func (ol OlValid) ReviewerName() string {
	return "ol/valid"
}

func (ol OlValid) Accepts(filePath string) bool {
	return true
}

func (ol OlValid) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return result, err
	}

	children := doc.Find("ol").Children()
	elem := children.Not("li")
	for i := 0; i < elem.Length(); i++ {
		result = append(result, Fault{
			Reviewer: ol.ReviewerName(),
			Line:     i + 1,
			Path:     path,
			Rule:     Rules["0008"],
		})
	}
	return result, nil
}
