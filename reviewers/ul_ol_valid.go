package reviewers

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type OlUlValid struct{}

func (ol OlUlValid) ReviewerName() string {
	return "ol-ul/valid"
}

func (ol OlUlValid) Accepts(filePath string) bool {
	return true
}

func (ol OlUlValid) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return result, err
	}

	olulSelection := doc.Find("ol").AddSelection(doc.Find("ul"))
	notli := olulSelection.Children().Not("li")

	for _, n := range notli.Nodes {
		result = append(result, Fault{
			Reviewer: ol.ReviewerName(),
			Line:     n.Line,
			Path:     path,
			Rule:     Rules["0008"],
		})
	}
	return result, nil
}
